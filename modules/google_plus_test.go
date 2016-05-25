package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGp(t *testing.T) {
	module := Gp{&Stats{}}
	module.Url = "http://yandex.ru"
	stats := module.CalculateShares()

	assert.NotEmpty(t, stats.Shares)
}