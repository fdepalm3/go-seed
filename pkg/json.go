package pkg

import (
	"encoding/json"
	"os"
)

func MustMarshal(elem interface{}) string {
	if res, err := json.Marshal(elem); err != nil {
		panic(err)
	} else {
		return string(res)
	}
}

func ReadJsonFile(filename string, reader any) error {
	jsonConfig, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jsonConfig.Close()

	jsonParser := json.NewDecoder(jsonConfig)
	if err = jsonParser.Decode(reader); err != nil {
		return err
	}

	return nil
}
