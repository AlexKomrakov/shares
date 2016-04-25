package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alexkomrakov/shares/src"
)

func TestGp(t *testing.T) {
	module := Gp{&shares.Stats{}}
	module.Url = "http://yandex.ru"
	stats := module.CalculateShares()

	assert.NotEmpty(t, stats.Shares)
}