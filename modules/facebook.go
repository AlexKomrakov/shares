package modules

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type FacebookResponse struct {
	ShareCount int `json:"share_count"`
}

type Facebook struct {
	*Stats
}

func (f Facebook) SetUrl(url string) {
	f.GetStats().Url = url
}

func (f Facebook) GetStats() *Stats {
	return f.Stats
}

func (m Facebook) GetShares() int {
	return m.GetStats().Shares
}

func (facebook Facebook) CalculateShares() *Stats {
	u, err := url.Parse(facebook.Url)
	if err != nil {
		panic(err)
	}

	response, err := http.Get("https://api.facebook.com/method/links.getStats?urls=" + u.String() + "&format=json")
	if err != nil {
		panic(err)
	}

	response_struct := make([]FacebookResponse, 1)
	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(response_body, &response_struct)

	if err != nil {
		panic(err)
	}

	facebook.Shares = response_struct[0].ShareCount

	return facebook.Stats
}