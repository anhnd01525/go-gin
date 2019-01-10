package helper

import (
	"os"

	jsoniter "github.com/json-iterator/go"
)

// Hàm parse dữ liệu JSON

func DecodeDataFromJsonFile(f *os.File, data interface{}) error {
	jsonParser := jsoniter.NewDecoder(f)
	err := jsonParser.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}
