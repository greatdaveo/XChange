package main

import (
	"log"
	"net/http"
	"os"

	"github.com/greatdaveo/XChange/controllers"
	"github.com/greatdaveo/XChange/rates"
	"github.com/greatdaveo/XChange/routes"
	"github.com/greatdaveo/XChange/services"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	apiKey := os.Getenv("EXCHANGE_RATE_API_KEY")

	// To initialize the rate provider
	externalAPI := &rates.RateProviders{APIKey: apiKey}
	// To pass it to the RateService
	rateService := services.NewRateService([]rates.RateProviders{*externalAPI})
	// To inject RateService into controller
	r := routes.RegisterRoutes(controllers.ConvertCurrency(rateService))

	// port := ":8080"

	log.Println("📌 Server running on :8080")
	http.ListenAndServe(":8080", r)
}
