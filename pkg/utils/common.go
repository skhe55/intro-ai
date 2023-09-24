package utils

import (
	"encoding/json"

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
