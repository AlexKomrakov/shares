package modules

import (
	"net/url"
	"regexp"
	"net/http"
	"io/ioutil"
	"strconv"
)

type Ok struct {
	*Stats
}

func (m Ok) SetUrl(url string) {
	m.GetStats().Url = url
}

func (m Ok) GetStats() *Stats {
	return m.Stats
}

func (m Ok) GetShares() int {
	return m.GetStats().Shares
}

func (m Ok) CalculateShares() *Stats {
	u, err := url.Parse(m.Url)
	if err != nil {
		panic(err)
	}

	response, err := http.Get("https://connect.ok.ru/dk?st.cmd=extLike&uid=odklcnt0&ref=" + u.String())
	if err != nil {
		panic(err)
	}

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`,'(.+)'\)`)
	result := re.FindAllStringSubmatch(string(response_body), -1)
	if (len(result) > 0) {
		str_count := result[0][len(result[0])-1]
		m.Shares, _ = strconv.Atoi(str_count)
	}

	return m.Stats
}