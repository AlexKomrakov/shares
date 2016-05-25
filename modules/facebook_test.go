package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFacebook(t *testing.T) {
	module := Facebook{&Stats{}}
	module.SetUrl("https://meduza.io/news/2016/03/30/v-windows-poyavitsya-bash")
	stats := module.CalculateShares()

	assert.NotEmpty(t, stats.Shares)
}