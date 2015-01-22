package pages

import (
	"encoding/json"
	"io/ioutil"
)

type Config map[string]interface{}

func (c Config) Name() string {
	return c["name"].(string)
}

func (c Config) Author() string {
	return c["author"].(string)
}

func LoadConfig(path string) (Config, error) {
	var m map[string]interface{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)
	return m, err
}
