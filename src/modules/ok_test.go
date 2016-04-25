package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alexkomrakov/shares/src"
)

func TestOk(t *testing.T) {
	module := Ok{&shares.Stats{}}
	module.Url = "http://worldru.ru/index.php?nma=news&fla=stat&nums=48756"
	stats := module.CalculateShares()

	assert.NotEmpty(t, stats.Shares)
}