package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Threatminer_json struct {
	Status_code    string   `json:"Status_code"`
	Status_message string   `json:"Status_message"`
	Results        []string `json:"Results"`
}

func Threatminer_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	r, err := http.Get(fmt.Sprintf("https://api.threatminer.org/v2/domain.php?q=%s&rt=5", domain))

	if err != nil {
		fmt.Println(err)
	}

	url_byte, _ := ioutil.ReadAll(r.Body)

	var Src Threatminer_json

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src.Results {
		filename.WriteString(Src.Results[index] + "\n")
		fmt.Println(Src.Results[index])
	}
}
