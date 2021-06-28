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

type Trails_json struct {
	Subdomains      []string `json:"Subdomains"`
	Subdomain_count int      `json:"Subdomain_count"`
	Meta            []string `json:"Meta"`
	Endpoint        string   `json:"Endpoint"`
}

func Security_trails_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	url := fmt.Sprintf("https://api.securitytrails.com/v1/domain/%s/subdomains", domain)

	r, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	r.Header.Add("Accept", "application/json")

	r.Header.Add("APIKEY", configuration.Security_trails_api)

	q := r.URL.Query()

	q.Add("children_only", "false")

	res, _ := http.DefaultClient.Do(r)

	url_byte, _ := ioutil.ReadAll(res.Body)

	var Src Trails_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src.Subdomains {
		filename.WriteString(fmt.Sprintf("%s%s%s%s", Src.Subdomains[index], ".", domain, "\n"))
		fmt.Printf("%s%s%s%s", Src.Subdomains[index], ".", domain, "\n")
	}
}
