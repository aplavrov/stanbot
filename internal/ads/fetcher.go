package ads

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Fetcher struct {
}

func New() *Fetcher {
	return &Fetcher{}
}

const (
	cityExpertURL = "https://cityexpert.rs/api/Search?req={\"cityId\":1,\"rentOrSale\":\"r\",\"searchSource\":\"regular\",\"sort\":\"datedsc\"}"
)

func (f *Fetcher) GetAll() []Property {
	client := http.Client{Timeout: 3 * time.Second}

	req, _ := http.NewRequest(http.MethodGet, cityExpertURL, nil)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		log.Fatal("failed to unmarshal", err)
	}
	return res.Result
}

func (f *Fetcher) GetLast() Property {
	client := http.Client{Timeout: 3 * time.Second}

	req, _ := http.NewRequest(http.MethodGet, cityExpertURL, nil)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		log.Fatal("failed to unmarshal", err)
	}
	return res.Result[0]
}
