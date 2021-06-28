package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Crtsh_json []Domain1

type Domain1 struct {
	Issuer_ca_id    int    `json:"Issuer_ca_id"`
	Issuer_name     string `json:"Issuer_name"`
	Common_name     string `json:"Common_name"`
	Name_value      string `json:"Name_value"`
	Id              int    `json:"Id"`
	Entry_timestamp string `json:"Entry_timestamp"`
	Not_before      string `json:"Not_before"`
	Not_after       string `json:"Not_after"`
	Serial_number   string `json:"Serial_number"`
}

func Crtsh_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	r, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain))

	if err != nil {
		fmt.Println(err)
	}

	url_byte, _ := ioutil.ReadAll(r.Body)

	var Src Crtsh_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src {
		filename.WriteString(Src[index].Name_value + "\n")
		fmt.Println(Src[index].Name_value)
	}
}
