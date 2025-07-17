package rates

type RateProvider interface {
	FetchRate(from string, to string) (float64, error)
}

