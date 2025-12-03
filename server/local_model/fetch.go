package local_model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type LocalModel struct {
	Content    string `json:"content"`
	Prediction string `json:"prediction"`
	Label      int    `json:"label"`
}

func Fetch(text string) string {
	main := log.New(os.Stdout, "[LOCAL MODEL]: ", log.Ldate|log.Ltime|log.Lshortfile)

	main.Println("fetching to local model....")
	encoded := url.QueryEscape(text)
	urlStr := fmt.Sprintf("%s/text?text=%s", os.Getenv("LOCAL_MODEL_URL"), encoded)
	response, err := http.Get(urlStr)

	if err != nil {
		main.Printf("http get error: %v", err)
		return "-1"
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		main.Printf("read body error: %v", err)
		return "-1"
	}

	main.Println("extracting the body mate")

	if response.StatusCode != http.StatusOK {
		main.Printf("unexpected status %s from local model, body: %s", response.Status, string(body))
		return "-1"
	}

	var PDIDDY LocalModel // i ran out of name to name this var soooooooooooo
	err = json.Unmarshal(body, &PDIDDY)

	if err != nil {
		main.Printf("json unmarshal error: %v", err)
		main.Printf("body was: %s", string(body))
		return "-1"
	}

	main.Println("done job")

	return strconv.Itoa(PDIDDY.Label) // 1 ands 0
}
