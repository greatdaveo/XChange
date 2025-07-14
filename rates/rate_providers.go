package rates

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type RateProviders struct {
	APIKey string
}

func (e *RateProviders) FetchRate(from, to string) (float64, error) {
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", e.APIKey, from)

	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	var data struct {
		Result          string             `json:"result"`
		ConversionRates map[string]float64 `json:"conversion_rates"`
	}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return 0, err
	}
	if data.Result != "success" {
		return 0, errors.New("API response failed")
	}

	// fmt.Print("Response:", data)

	// fmt.Println("Available rates:")
	// for k := range data.ConversionRates {
	// 	fmt.Print(k + " ")
	// }

	rate, ok := data.ConversionRates[to]
	if !ok {
		return 0, errors.New("❌ target currency not found")
	}

	fmt.Println("📌 Final rate returned from API:", rate)

	return rate, nil
}
