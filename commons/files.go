package commons

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadFromJson(input string, data interface{}) error {
	return json.Unmarshal([]byte(input), &data)
}

func ToJson(data interface{}) {
	res, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(res))
}

func ToJsonFile(filepath string, data interface{}) {
	content, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile(filepath, content, 0644)
}
