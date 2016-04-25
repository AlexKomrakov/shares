package modules

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alexkomrakov/shares/src"
)

func TestMm(t *testing.T) {
	module := Mm{&shares.Stats{}}
	module.Url = "http://yandex.ru"
	stats := module.CalculateShares()

	assert.NotEmpty(t, stats.Shares)
}