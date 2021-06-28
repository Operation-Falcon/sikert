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

type Virus_total_json struct {
	BitDefendercategory              string   `json:"BitDefendercategory"`
	Undetected_downloaded_samples    []string `json:"Undetected_downloaded_samples"`
	Whois_timestamp                  int      `json:"Whois_timestamp"`
	Detected_downloaded_samples      []string `json:"Detected_downloaded_samples"`
	Detected_referrer_samples        []string `json:"Detected_referrer_samples"`
	AlphaMountain_ai_category        string   `json:"AlphaMountain_ai_category"`
	Webutation_domain_info           []string `json:"Webutation_domain_info"`
	Sophos_category                  string   `json:"Sophos_category"`
	Comodo_Valkyrie_Verdict_category string   `json:"Comodo_Valkyrie_Verdict_category"`
	Subdomains                       []string `json:"Subdomains"`
	Undetected_referrer_samples      []string `json:"Undetected_referrer_samples"`
	Resolutions                      []string `json:"Resolutions"`
	Detected_communicating_samples   []string `json:"Detected_communicating_samples"`
	Domain_siblings                  []string `json:"Domain_siblings"`
	BitDefender_domain_info          string   `json:"BitDefender_domain_info"`
	Whois                            string   `json:"Whois"`
	Response_code                    string   `json:"Response_code"`
	Forcepoint_ThreatSeeker_category string   `json:"Forcepoint_ThreatSeeker_category"`
	Verbose_msg                      string   `json:"Verbose_msg"`
	Detected_urls                    []string `json:"Detected_urls"`
	Undetected_communicating_samples []string `json:"Undetected_communicating_samples"`
	Undetected_urls                  []string `json:"undetected_urls"`
}

func Virus_total_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	url := "https://www.virustotal.com/vtapi/v2/domain/report"

	r, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	q := r.URL.Query()
	q.Add("apikey", configuration.Virus_total_api)
	q.Add("domain", domain)

	res, _ := http.DefaultClient.Do(r)

	url_byte, _ := ioutil.ReadAll(res.Body)

	var Src Virus_total_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src.Subdomains {
		filename.WriteString(Src.Subdomains[index] + "\n")
		fmt.Println(Src.Subdomains[index])
	}
}
