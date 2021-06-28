package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sikert/configuration"
	"sync"
)

type Binary_json struct {
	Query    string   `json:"Query"`
	Page     int      `json:"Page"`
	Pagesize int      `json:"Pagesize"`
	Total    int      `json:"Total"`
	Events   []string `json:"Events"`
}

func Binary_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	url := fmt.Sprintf("https://api.binaryedge.io/v2/query/domains/subdomain/%s", domain)

	r, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	r.Header.Add("Content-Type", "application/json")

	r.Header.Add("X-Key", configuration.Binary_api)

	res, _ := http.DefaultClient.Do(r)

	url_byte, _ := ioutil.ReadAll(res.Body)

	var Src Binary_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src.Events {
		filename.WriteString(Src.Events[index] + "\n")
		fmt.Printf("%s%s", Src.Events[index], "\n")
	}
}
