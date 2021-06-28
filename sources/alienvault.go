package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Alienvault_json struct {
	Passive_dns []Domain `json:"Passive_dns"`
	Count       int      `json:"Count"`
}

type Domain struct {
	Address        string `json:"Address"`
	First          string `json:"First"`
	Last           string `json:"Last"`
	Hostname       string `json:"Hostname"`
	Record_type    string `json:"Record_type"`
	Indicator_link string `json:"Indicator_link"`
	Flag_url       string `json:"Flag_url"`
	Flag_title     string `json:"Flag_title"`
	Asset_type     string `json:"Asset_type"`
	Asn            string `json:"Asn"`
}

func Alienvault_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	r, err := http.Get(fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/domain/%s/passive_dns", domain))

	if err != nil {
		fmt.Println(err)
	}

	url_byte, _ := ioutil.ReadAll(r.Body)

	var Src1 Alienvault_json

	json.Unmarshal(url_byte, &Src1)

	filename, _ := os.Create(file)

	for index, _ := range Src1.Passive_dns {

		filename.WriteString(Src1.Passive_dns[index].Hostname + "\n")
		fmt.Printf("%s%s", Src1.Passive_dns[index].Hostname, "\n")
	}

}
