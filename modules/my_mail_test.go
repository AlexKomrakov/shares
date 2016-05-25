package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMm(t *testing.T) {
	module := Mm{&Stats{}}
	module.Url = "http://yandex.ru"
	stats := module.CalculateShares()

	assert.NotEmpty(t, stats.Shares)
}