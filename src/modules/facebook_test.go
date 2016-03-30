package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetStats(t *testing.T) {
	module := Facebook{}
	module.Url = "https://meduza.io/news/2016/03/30/v-windows-poyavitsya-bash"
	stats := module.GetStats()

	assert.NotEmpty(t, stats.Shares)
}