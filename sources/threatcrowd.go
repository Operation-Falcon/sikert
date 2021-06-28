package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Threatcrowd_json struct {
	Response_code string   `json:"Response_code"`
	Resolutions   []string `json:"Resolutions"`
	Hashes        []string `json:"Hashes"`
	Emails        []string `json:"Emails"`
	Subdomains    []string `json:"Subdomains"`
	References    []string `json:"References"`
	Votes         int      `json:"Votes"`
	Permalink     []string `json:"Permalink"`
}

func Threatcrowd_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	r, err := http.Get(fmt.Sprintf("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s", domain))

	if err != nil {
		fmt.Println(err)
	}

	url_byte, _ := ioutil.ReadAll(r.Body)

	var Src Threatcrowd_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src.Subdomains {
		filename.WriteString(Src.Subdomains[index] + "\n")
		fmt.Println(Src.Subdomains[index])
	}
}
