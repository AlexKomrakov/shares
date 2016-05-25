package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	module := Ok{&Stats{}}
	module.Url = "http://worldru.ru/index.php?nma=news&fla=stat&nums=48756"
	stats := module.CalculateShares()

	assert.NotEmpty(t, stats.Shares)
}