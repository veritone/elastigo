// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elastigo

import (
	"encoding/json"
	"flag"
	"log"
)

var (
	_                = log.Ldate
	eshost   *string = flag.String("host", "localhost", "Elasticsearch Server Host Address")
	esport   *string = flag.String("port", "9200", "Elasticsearch Server Host Port")
	logLevel *string = flag.String("logging", "info", "Which log level: [debug,info,warn,error,fatal]")
)

func GetJson(input interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	bytes, err := json.Marshal(input)

	if err == nil {
		err = json.Unmarshal(bytes, &result)
	}

	return result, err
}

func HasKey(input map[string]interface{}, key string) bool {
	if _, ok := input[key]; ok {
		return true
	}

	return false
}
