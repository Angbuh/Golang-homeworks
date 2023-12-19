package reader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/angbuh/golang-homeworks/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return readJSONFromReader(f)
}

func readJSONFromReader(r io.Reader) ([]types.Employee, error) {
	var data []types.Employee

	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
