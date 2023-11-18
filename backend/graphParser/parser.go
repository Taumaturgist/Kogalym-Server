package main

import (
	"bufio"
	"fmt"
	"kogalym-backend/helpers"
	"os"
	"strings"
)

type graph struct {
	name        string
	units       string
	description string
	data        map[string]string
}

func main() {
	parseFiles("./files")
}

func parse(filepath string) {
	startGraphTypesWriteString := "DEPTH   .M                                                  :DEPTH"
	endGraphTypesWriteString := "~Other information"
	startGraphDataString := "~ASCII Log Data"

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	var words []string

	isStartedGraphTypes, isEndedGraphTypes, isStartedData := false, false, false
	scanner := bufio.NewScanner(f)
	var graphs = map[string]graph{}
	var graphsNames []string
	for scanner.Scan() {
		line := string(scanner.Bytes())

		if !isStartedGraphTypes && line == startGraphTypesWriteString {
			isStartedGraphTypes = true
			continue
		}

		if isStartedGraphTypes && !isEndedGraphTypes {
			if !isEndedGraphTypes && line == endGraphTypesWriteString {
				isEndedGraphTypes = true
				continue
			}

			graphName, after, _ := strings.Cut(line, " ")
			units, after, _ := strings.Cut(strings.TrimSpace(after), " ")
			description, after, _ := strings.Cut(strings.TrimSpace(after), " ")

			graphsNames = append(graphsNames, graphName)
			graphs[graphName] = graph{
				name:        graphName,
				units:       units,
				description: description,
				data:        make(map[string]string),
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
			graphs[name].data[depth] = chunk[idx]
		}
	}

	// todo запись в БД

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func parseFiles(dir string) {
	entries, err := os.ReadDir(dir)
	helpers.CheckErr(err)

	for _, e := range entries {
		if e.IsDir() {
			parseFiles(dir + "/" + e.Name())
		} else {
			parse(dir + "/" + e.Name())
		}
	}
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
