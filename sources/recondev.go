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

type Recondev_json []Doomsday

type Doomsday struct {
	Domains    []string `json:"Domains"`
	IP         string   `json:"IP"`
	Rawdomains []string `json:"Rawdomains"`
	Rawport    string   `json:"Rawport"`
	Rawip      string   `json:"Rawip"`
}

func Recondev_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	url := fmt.Sprintf("https://recon.dev/api/search?key=%s&domain=%s", configuration.Recondev_api, domain)

	r, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	r.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(r)

	url_byte, _ := ioutil.ReadAll(res.Body)

	var Src Recondev_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src {
		for dom, _ := range Src[index].Rawdomains {
			filename.WriteString(Src[index].Rawdomains[dom] + "\n")
			fmt.Println(Src[index].Rawdomains[dom])
		}
	}
}
