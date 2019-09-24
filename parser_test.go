package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMerge(t *testing.T) {
	json1 := `{
				"key": "/node/nm/device",
				"value": {
					"fake": {
						"key": "nothing"
					},
					"monitor": {
        				"cpu": 29,
						"hi": {
						"key": "haha"
						}
        			}
     			}
 			}`

	json2 := `{
	  			"value": {
        			"gateway": {
        				"ip": "196.12.22.33",
          				"port": 1234,
						"hello": {
							"key": "hello"
						}
        			},
        			"monitor": {
        				"cpu": 30,
						"hi": {
							"key": "hi"
						}
        			}
     			}
 			}`

	var d1, d2 map[string]interface{}
	_ = json.Unmarshal([]byte(json1), &d1)
	_ = json.Unmarshal([]byte(json2), &d2)

	data, _ := json.Marshal(JsonMerge(d1, d2))
	fmt.Println(string(data))
}
