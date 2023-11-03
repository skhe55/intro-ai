package utils

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

func ToJSON[T any](v T) ([]byte, error) {
	result, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ParseJSON[T any](src string) (T, error) {
	var args T

	if err := json.Unmarshal([]byte(src), &args); err != nil {
		return *(new(T)), err
	}

	validator := validator.New()
	if err := validator.Struct(args); err != nil {
		return *(new(T)), err
	}

	return args, nil
}

func ConvertToValidMimeType(mimeType string) (string, error) {
	if mimeType == "image/jpeg" {
		return "jpg", nil
	} else if mimeType == "image/png" {
		return "png", nil
	}

	return "", errors.New("not recognized mime type")
}

func ConvertStringToFloat64SliceOfSlices(v string) [][]float64 {
	var result [][]float64
	v = strings.Trim(v, "{}")
	pairs := strings.Split(v, "},{")

	for _, pair := range pairs {
		pair = strings.Trim(pair, "{}")
		nums := strings.Split(pair, ",")
		var subResult []float64
		for _, num := range nums {
			f, _ := strconv.ParseFloat(num, 64)
			subResult = append(subResult, f)
		}
		result = append(result, subResult)
	}

	return result
}
