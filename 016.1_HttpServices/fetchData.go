package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Quotes struct {
	Quote string
}

func FetchTest() {
	url := "https://api.kanye.rest/"
	r, err := http.NewRequest("GET", url, nil)
	client := http.Client{Timeout: time.Second * 5}
	res, err := client.Do(r)

	defer func() {
		if err != nil {
			fmt.Println(err)
		}
		res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)

	quote := Quotes{}
	json.Unmarshal(body, &quote)

	fmt.Println(quote.Quote)
}
