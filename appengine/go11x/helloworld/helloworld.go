// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample helloworld is a basic App Engine flexible app.
//modified to display random line from a file
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type Quotes struct{
	Quotes []Quote `json:"quotes"`
}

type Quote struct{
	Quote string `json:"quote"`
}

type BotResponses struct{
	BotResponses []Quote `json:"responses"`
}

func main() {
	http.HandleFunc("/", handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	sender := r.URL.Query().Get("sender")
	target := r.URL.Query().Get("target")

	if target == "Anthony" || "Xander" || "@XiJaroAndPitch" || "@xijaroandpitch"{

		f, _ := ioutil.ReadFile("random-makeout-quotes.json")
		data := Quotes{}
	
		_ = json.Unmarshal([]byte(f), &data)
	
		i := rand.Intn(len(data.Quotes))
	
		line := data.Quotes[i].Quote

		if sender != "" {
			line = strings.Replace(line, "$(sender)", sender, -1)
		}
	
		if target != "" {
			line = strings.Replace(line, "$(target)", target, -1)
		}
	
		fmt.Fprint(w, line)	
	}
	else{
		
		f, _ := ioutil.ReadFile("random-makeout-responses.json")
		data := BotResponses{}

		_ = json.Unmarshal([]byte(f), &data)

		i := rand.Intn(len(data.BotResponses))

		line := data.BotResponses[i].Quote

		if sender != "" {
			line = strings.Replace(line, "$(sender)", sender, -1)
		}
	
		if target != "" {
			line = strings.Replace(line, "$(target)", target, -1)
		}
	
		fmt.Fprint(w, line)
	}
}
