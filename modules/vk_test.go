package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVk(t *testing.T) {
	module := Vk{&Stats{}}
	module.Url = "https://meduza.io/news/2016/03/30/v-windows-poyavitsya-bash"
	stats := module.CalculateShares()

	assert.NotEmpty(t, stats.Shares)
}