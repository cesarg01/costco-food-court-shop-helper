package main

import (
	"net/http"
	"strconv"
)

// ShopDays struct containing the number of shop days.
type ShopDays struct {
	NumberOfShopDays int
}

// IndexPage ...
func IndexPage(w http.ResponseWriter, r *http.Request) {

	// Convert string to number for each form value from index.html
	totalDays := r.FormValue("shop_days")
	totalShopDays, _ := strconv.Atoi(totalDays)

	// Add the totalShopDays to the ShopDays struct
	days := ShopDays{
		NumberOfShopDays: totalShopDays,
	}

	renderInitialPage(w, "templates/index.html", days)
}
