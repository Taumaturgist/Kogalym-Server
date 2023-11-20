package graphParser

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"kogalym-backend/helpers"
	"kogalym-backend/models"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ParseData() {
	currentDir, err := os.Getwd()
	helpers.CheckErr(err)

	dir := currentDir + "/graphParser/files"

	entries, err := os.ReadDir(dir)
	helpers.CheckErr(err)

	for _, e := range entries {
		if e.IsDir() {
			parseFiles(dir+"/"+e.Name(), e.Name())
		}
	}
}

func parse(dir string, filename string, mineName string, isPriorityFile bool) {
	startGraphTypesWriteString := "DEPTH   .M                                                  :DEPTH"
	endGraphTypesWriteString := "~Other information"
	startGraphDataString := "~ASCII Log Data"
	startDepthDescription := ":START DEPTH"
	stopDepthDescription := ":STOP DEPTH"
	stepDepthDescription := ":STEP"

	filePath := ""
	if isPriorityFile {
		filePath = dir + "/*/" + filename
	} else {
		filePath = dir + "/" + filename
	}
	fmt.Println("++++++++++++++++++++++++++")
	fmt.Println(filePath)

	f, err := os.Open(filePath)
	helpers.CheckErr(err)
	defer f.Close()

	var words []string

	isStartedGraphTypes, isEndedGraphTypes, isStartedData := false, false, false
	charMapDecoder := charmap.Windows1251.NewDecoder()
	reader := transform.NewReader(f, charMapDecoder)
	scanner := bufio.NewScanner(reader)
	var graphs = map[string]models.Graph{}
	var graphsNames []string
	var startDepth, stopDepth, stepDepth float64
	for scanner.Scan() {
		line := string(scanner.Bytes())

		if !isStartedGraphTypes && strings.Contains(line, startDepthDescription) {
			startDepth = getValue(line, startDepthDescription)
		}

		if !isStartedGraphTypes && strings.Contains(line, stopDepthDescription) {
			stopDepth = getValue(line, stopDepthDescription)
		}

		if !isStartedGraphTypes && strings.Contains(line, stepDepthDescription) {
			stepDepth = getValue(line, stepDepthDescription)
		}

		if !isStartedGraphTypes && line == startGraphTypesWriteString {
			isStartedGraphTypes = true
			continue
		}

		if isStartedGraphTypes && !isEndedGraphTypes {
			if !isEndedGraphTypes && line == endGraphTypesWriteString {
				isEndedGraphTypes = true
				continue
			}

			stringSlice := strings.Fields(line)

			graphName, units, description := stringSlice[0], stringSlice[1], stringSlice[2]

			graphsNames = append(graphsNames, graphName)
			graphs[graphName] = models.Graph{
				MineName:    mineName,
				Name:        graphName,
				Units:       strings.TrimPrefix(units, `.`),
				Description: strings.TrimPrefix(description, `:`),
				StartDepth:  math.Min(startDepth, stopDepth),
				StopDepth:   math.Max(stopDepth, startDepth),
				StepDepth:   math.Abs(stepDepth),
				Data:        make(map[float64]float64),
			}

			continue
		}

		if !isStartedData && line == startGraphDataString {
			isStartedData = true
			continue
		}

		if isStartedData {
			words = append(words, strings.Fields(line)...)
		}
	}

	wordsChunks := chunkSlice(words, len(graphsNames)+1)

	for _, chunk := range wordsChunks {
		depth := chunk[0]
		chunk = chunk[1:]
		for idx, name := range graphsNames {
			graphs[name].Data[depth] = chunk[idx]
		}
	}

	models.UpdateOrCreateGraphData(graphs, isPriorityFile)

	helpers.CheckErr(scanner.Err())
}

func parseFiles(dir string, mineName string) {
	entries, err := os.ReadDir(dir)
	helpers.CheckErr(err)

	// сначала обраабатываем файлы
	for _, e := range entries {
		if !e.IsDir() && strings.ToLower(filepath.Ext(dir+"/"+e.Name())) == `.las` {
			parse(dir, e.Name(), mineName, false)
		}
	}

	// потом обрабатываем файлы из вложенной папки *, как приоритетные
	for _, e := range entries {
		if e.IsDir() && e.Name() == `*` {
			entries, err := os.ReadDir(dir + "/*")
			helpers.CheckErr(err)
			for _, e := range entries {
				if !e.IsDir() && strings.ToLower(filepath.Ext(dir+"/*/"+e.Name())) == `.las` {
					parse(dir, e.Name(), mineName, true)
				}
			}
		}
	}
}

func getValue(line string, valueDescription string) float64 {
	startDepth, _, _ := strings.Cut(line, valueDescription)
	sl := strings.Split(startDepth, " ")
	number := sl[len(sl)-1]
	return parseFloat(number)
}

func chunkSlice(slice []string, chunkSize int) [][]float64 {
	var chunks [][]float64
	var floatSlice []float64
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		for i := 0; i < chunkSize; i++ {
			floatSlice = append(floatSlice, parseFloat(slice[i]))
		}

		chunks = append(chunks, floatSlice)
		slice = slice[chunkSize:]
	}

	return chunks
}

func parseFloat(number string) float64 {
	s, _ := strconv.ParseFloat(number, 32)

	return s
}
