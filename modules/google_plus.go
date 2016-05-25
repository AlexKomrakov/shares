package modules

import (
	"net/url"
	"regexp"
	"net/http"
	"io/ioutil"
	"strconv"
)

type Gp struct {
	*Stats
}

func (m Gp) SetUrl(url string) {
	m.GetStats().Url = url
}

func (m Gp) GetStats() *Stats {
	return m.Stats
}

func (m Gp) GetShares() int {
	return m.GetStats().Shares
}

func (m Gp) CalculateShares() *Stats {
	u, err := url.Parse(m.Url)
	if err != nil {
		panic(err)
	}

	response, err := http.Get("https://plusone.google.com/u/0/_/+1/fastbutton?count=true&url=" + u.String())
	if err != nil {
		panic(err)
	}

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`Oy">>?([\d]+)<`)
	result := re.FindAllStringSubmatch(string(response_body), -1)
	if (len(result) > 0) {
		str_count := result[0][len(result[0])-1]
		m.Shares, _ = strconv.Atoi(str_count)
	}

	return m.Stats
}