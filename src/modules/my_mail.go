package modules

import (
	"github.com/alexkomrakov/shares/src"
	"net/url"
	"regexp"
	"net/http"
	"io/ioutil"
	"strconv"
)

type Mm struct {
	*shares.Stats
}

func (m Mm) SetUrl(url string) {
	m.GetStats().Url = url
}

func (m Mm) GetStats() *shares.Stats {
	return m.Stats
}

func (m Mm) GetShares() int {
	return m.GetStats().Shares
}

func (m Mm) CalculateShares() *shares.Stats {
	u, err := url.Parse(m.Url)
	if err != nil {
		panic(err)
	}

	response, err := http.Get("http://connect.mail.ru/share_count?url_list=" + u.String())
	if err != nil {
		panic(err)
	}

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`shares":([\d]+),`)
	result := re.FindAllStringSubmatch(string(response_body), -1)
	if (len(result) > 0) {
		str_count := result[0][len(result[0])-1]
		m.Shares, _ = strconv.Atoi(str_count)
	}

	return m.Stats
}