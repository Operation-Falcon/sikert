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

type Chaos_json struct {
	Domain     string   `json:"Domain"`
	Subdomains []string `json:"Subdomains"`
	Count      int      `json:"Count"`
}

func Chaos_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	url := fmt.Sprintf("https://dns.projectdiscovery.io/dns/%s/subdomains", domain)

	r, err := http.NewRequest("GET", url, nil)

	r.Header.Add("Authorization", configuration.Chaos_api)

	if err != nil {
		fmt.Println(err)
	}

	res, _ := http.DefaultClient.Do(r)

	url_byte, _ := ioutil.ReadAll(res.Body)

	var Src Chaos_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src.Subdomains {
		filename.WriteString(fmt.Sprintf("%s%s%s%s", Src.Subdomains[index], ".", domain, "\n"))
		fmt.Printf("%s%s%s%s", Src.Subdomains[index], ".", domain, "\n")
	}
}
