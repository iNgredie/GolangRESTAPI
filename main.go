package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Funds struct {
	Id               int             `json:"id"`
	Name             string          `json:"name"`
	Ticker           string          `json:"ticker"`
	Amount           int64           `json:"amount"`
	PricePerItem     decimal.Decimal `json:"priceperitem"`
	PurchasePrice    decimal.Decimal `json:"purchaseprice"`
	PriceCurrent     decimal.Decimal `json:"pricecurrent"`
	PercentChanges   decimal.Decimal `json:"percentchanges"`
	YearlyInvestment decimal.Decimal `json:"yearlyinvestment"`
	ClearMoney       decimal.Decimal `json:"clearmoney"`
	DatePurchase     time.Time       `json:"datepurchase"`
	DateLastUpdate   time.Time       `json:"datelastupdate"`
	Type             string          `json:"type"`
}

func main() {
	fmt.Println("My REST server")

	r := mux.NewRouter()

	// fund
	// GET
	r.HandleFunc("/funds/usd/shares", getUSDFundsShares)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getUSDFundsShares(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var ArrShares = myCurrentFunds("share")
	json.NewEncoder(w).Encode(ArrShares)
}

func myCurrentFunds(fundType string) []Funds {
	var amountShares []Funds

	db, err := sql.Open("postgres", "postgres//postgres:1234@localhost/fin?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM fundsusd WHERE type = $1", fundType)

	for rows.Next() {
		f := Funds{}
		err = rows.Scan(&f.Id, &f.Name, &f.Ticker, &f.Amount, &f.PricePerItem, &f.PurchasePrice, &f.PriceCurrent, &f.PercentChanges, &f.YearlyInvestment, &f.ClearMoney, &f.DatePurchase, &f.DateLastUpdate, &f.Type)
		if err != nil {
			log.Fatal(err)
		}
		amountShares = append(amountShares, f)
	}
	return amountShares
}
