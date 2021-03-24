package mapping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Mapping struct {
	Request    Request  `json:"request"`
	Response   Response `json:"response"`
	RawContent string
}

type Request struct {
	Url    string
	Method string
}

type Response struct {
	Status int
	Body   string
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
