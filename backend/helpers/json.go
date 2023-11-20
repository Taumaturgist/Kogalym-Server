package helpers

import (
	"encoding/json"
	"sort"
)

func FromMap(data map[float64]float64) (string, error) {
	data = sortMap(data)
	// TODO ошибка записи json, потому что в json ключом может быть только строка
	// Преобразуем map в JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Преобразуем байты JSON в строку
	jsonString := string(jsonBytes)

	return jsonString, nil
}

func ToMap(rawData string) map[float64]float64 {
	// Создаем карту для хранения данных
	data := make(map[float64]float64)

	// Распаковываем JSON-строку в карту
	json.Unmarshal([]byte(rawData), &data)

	return data
}

func sortMap(data map[float64]float64) map[float64]float64 {
	var result map[float64]float64
	indexes := sortMapIndexes(data)
	for _, k := range indexes {
		result[k] = data[k]
	}

	return result
}

func sortMapIndexes(data map[float64]float64) []float64 {
	var indexes []float64
	for k, _ := range data {
		indexes = append(indexes, k)
	}
	sort.Float64s(indexes)

	return indexes
}
