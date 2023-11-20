package models

import (
	"fmt"
	"kogalym-backend/helpers"
	"math"
)

type Graph struct {
	MineName    string
	Name        string
	Units       string
	Description string
	StartDepth  float64
	StopDepth   float64
	StepDepth   float64
	Data        map[float64]float64
}

func UpdateOrCreateGraphData(graphs map[string]Graph, isPriority bool) {
	for _, graph := range graphs {
		fmt.Println(graph.Name)
		if isPriority {
			updateOldGraphDataIfExists(graph)
		} else {
			insertGraphData(graph)
		}
	}
}

func insertGraphData(graph Graph) {
	stmt, err := DB.Prepare(
		"INSERT INTO graph_data (mine_name, graph_name, units, description, start_depth, stop_depth, step_depth, data) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
	)
	helpers.CheckErr(err)

	data, err := helpers.FromMap(graph.Data)

	helpers.CheckErr(err)
	_, err = stmt.Exec(
		graph.MineName,
		graph.Name,
		graph.Units,
		graph.Description,
		graph.StartDepth,
		graph.StopDepth,
		graph.StepDepth,
		data,
	)
	helpers.CheckErr(err)
}

func updateOldGraphDataIfExists(graph Graph) {
	var rawData string
	var startDepth, stopDepth, stepDepth float64

	stmt, err := DB.Prepare(
		"SELECT start_depth, stop_depth, step_depth, data from graph_data WHERE graph_name = ? AND mine_name = ?;",
	)
	helpers.CheckErr(err)

	stmt.QueryRow(graph.Name, graph.MineName).Scan(&startDepth, &stopDepth, &stepDepth, &rawData)

	if rawData == "" {
		insertGraphData(graph)
		return
	}

	data := helpers.ToMap(rawData)
	for key, value := range graph.Data {
		data[key] = value
	}

	startDepth = math.Min(graph.StartDepth, startDepth)
	stopDepth = math.Max(graph.StopDepth, stopDepth)

	updateStmt, err := DB.Prepare(
		"UPDATE graph_data SET start_depth = ?, stop_depth = ?, data = ? WHERE mine_name = ? AND graph_name = ?;",
	)
	helpers.CheckErr(err)

	dataForInsert, err := helpers.FromMap(data)

	helpers.CheckErr(err)
	_, err = updateStmt.Exec(
		startDepth,
		stopDepth,
		dataForInsert,
		graph.MineName,
		graph.Name,
	)
	helpers.CheckErr(err)
}
