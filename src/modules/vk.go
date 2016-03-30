package modules

import "github.com/alexkomrakov/shares/src"

type Vk struct {
	shares.Stats
}

func (vk Vk) GetStats() shares.Stats {

	return vk.Stats
}