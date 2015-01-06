// To make use of some of our collections knowledge, we will parse some JSON
// configuration into our application.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// we will create a function called LoadConfig that will load up our JSON
// configuration file
func LoadConfig(path string) (map[string]interface{}, error) {
	var m map[string]interface{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)
	return m, err
}

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
