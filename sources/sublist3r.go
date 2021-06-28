package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func Sublist3r_enum(domain, file string, wg *sync.WaitGroup) {

	defer wg.Done()

	r, err := http.Get(fmt.Sprintf("https://api.sublist3r.com/search.php?domain=%s", domain))

	if err != nil {
		fmt.Println(err)
	}

	url_byte, _ := ioutil.ReadAll(r.Body)

	var Src []string

	json.Unmarshal(url_byte, &Src)

	filename, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)

	for index, _ := range Src {
		filename.WriteString(Src[index] + "\n")
		fmt.Println(Src[index])
	}
}
