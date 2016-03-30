package modules

import (
	"github.com/alexkomrakov/shares/src"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type FacebookResponse struct {
	ShareCount int `json:"share_count"`
}

type Facebook struct {
	shares.Stats
}

func (facebook Facebook) GetStats() shares.Stats {
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