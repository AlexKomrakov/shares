package shares

type Stats struct {
	Url string
	Shares int
}

type HasStats interface {
	GetStats() *Stats
	GetShares() int
	CalculateShares() *Stats
	SetUrl(string)
}