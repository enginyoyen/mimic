package mapping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Mapping struct {
	Request    Request  `json:"request"`
	Response   Response `json:"response"`
	RawContent string
}

type Request struct {
	Url     string
	Method  string
	Headers map[string]string
}

type Response struct {
	Status    int
	Body      json.RawMessage
	WithDelay time.Duration
	Headers   map[string]string
}

func MapRequests(files *[]string) *[]Mapping {
	var mappings []Mapping
	if len(*files) == 0 {
		return &mappings
	}
	for _, file := range *files {

		jsonFile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var mapping Mapping
		json.Unmarshal(byteValue, &mapping)
		mapping.RawContent = string(byteValue)
		mappings = append(mappings, mapping)
		fmt.Printf("%+v\n", mapping)
	}
	return &mappings
}
