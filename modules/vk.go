package modules

import (
	"net/url"
	"regexp"
	"net/http"
	"io/ioutil"
	"strconv"
)

type Vk struct {
	*Stats
}

func (v Vk) SetUrl(url string) {
	v.GetStats().Url = url
}

func (v Vk) GetStats() *Stats {
	return v.Stats
}

func (m Vk) GetShares() int {
	return m.GetStats().Shares
}

func (vk Vk) CalculateShares() *Stats {
	u, err := url.Parse(vk.Url)
	if err != nil {
		panic(err)
	}

	response, err := http.Get("http://vk.com/share.php?act=count&url=" + u.String())
	if err != nil {
		panic(err)
	}

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`, (.+)\)`)
	result := re.FindAllStringSubmatch(string(response_body), -1)
	if (len(result) > 0) {
		str_count := result[0][len(result[0])-1]
		vk.Shares, _ = strconv.Atoi(str_count)
	}

	return vk.Stats
}