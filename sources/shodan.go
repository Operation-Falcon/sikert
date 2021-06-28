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

type Shodan_json struct {
	Domain     string   `json:"Domains"`
	Tags       []string `json:"Tags"`
	Data       []string `json:"Data"`
	Subdomains []string `json:"Subdomains"`
	More       bool     `json:"More"`
}

func Shodan_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	r, err := http.Get(fmt.Sprintf("https://api.shodan.io/dns/domain/%s?key=%s", domain, configuration.Shodan_api))

	if err != nil {
		fmt.Println(err)
	}
	url_byte, _ := ioutil.ReadAll(r.Body)

	var Src Shodan_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src.Subdomains {
		filename.WriteString(fmt.Sprintf("%s%s%s%s", Src.Subdomains[index], ".", domain, "\n"))
		fmt.Printf("%s%s%s%s", Src.Subdomains[index], ".", domain, "\n")
	}
}
