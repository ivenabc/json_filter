//+build ignore
package json_filter

import (
	"testing"
	"fmt"
	"encoding/json"
)


var parseStr = `
{
  "baseKeys": ["hello"],
  "objectKeys":["another", "alsoInner"],
  "arrayKeys": ["inner"],
  "inner": {
    "baseKeys": [
      "value2"
    ],
    "objectKeys": ["tt"],
    "tt": {
      "baseKeys":["value5"]
    }
  },
  "alsoInner":{
    "baseKeys": ["value1"]
  },
  "another":{
    "baseKeys": ["value3"]
  }
}
`

var jsonStr = `{
  "inner":[{
    "value1":10,
    "value2":22,
    "tt":{
      "value4": 122,
      "value5": 112,
      "value6": 102
    },
    "value3":122
  }],
  "alsoInner":{
    "value1":20,
    "value2":42,
    "value3":92
  },
  "hello": 11,
  "world": 13,
  "another":{
    "value1":null,
    "value2":null,
    "value3":null
  }
}
`


func TestGenerateFilterJson(t *testing.T) {
	i := GenerateFilterJson([]byte(jsonStr), []byte(parseStr))

	data, _ := json.Marshal(i)
	fmt.Println(string(data))
}