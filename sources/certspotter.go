package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Certspotter_json []struct {
	Id            int      `json:"Id"`
	Tbs_sha256    []string `json:"Tbs_sha256"`
	Dns_names     []string `json:"Dns_names"`
	Pubkey_sha256 []string `json:"Pubkey_sha256"`
	Not_before    string   `json:"Not_before"`
	Not_after     string   `json:"Not_after"`
}

func Certspotter_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	r, err := http.Get(fmt.Sprintf("https://api.certspotter.com/v1/issuances?domain=%s&include_subdomains=true&expand=dns_names", domain))

	if err != nil {
		fmt.Println(err)
	}

	url_byte, _ := ioutil.ReadAll(r.Body)

	var Src Certspotter_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src {
		for dom, _ := range Src[index].Dns_names {
			filename.WriteString(Src[index].Dns_names[dom] + "\n")
			fmt.Printf("%s%s", Src[index].Dns_names[dom], "\n")
		}
	}
}
