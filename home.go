package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type ShopDetails struct {
	PizzaSold              int
	FeSmoothieSold         int
	KioskSmoothieSold      int
	FeColdBrewSold         int
	KioskColdBrewSold      int
	FeColdBrewLatteSold    int
	KioskColdBrewLatteSold int
	FeSodaSold             int
	KioskSodaSold          int
	IceCreamSold           int
	HotDogSold             int
	SipLidsToBuy           int
	SodaLidsToBuy          int
	SodaCupsToBuy          int
	IceCreamLidsToBuy      int
	SixteenOunceCupsToBuy  int
	PaperStrawsToBuy       int
}

// IndexPage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	//var pizzaSold string
	//if r.Method != http.MethodPost {

	//	return
	//}

	steel := Steel{}
	// Smoothie Base
	steel.SmoothieBase.Quantity = 2
	steel.SmoothieBase.One6OzGallon = 8
	// Dairy Mix
	steel.DairyMix.Quantity = 10
	steel.DairyMix.One6OzGallon = 8
	// Soda Cups
	steel.SodaCups.Quantity = 1000
	steel.SodaCups.Sleeve = 50
	steel.SodaCups.Sleeves = 20
	// 16 oz Cups
	steel.One6OzCups.Quantity = 1000
	steel.One6OzCups.Sleeve = 50
	steel.One6OzCups.Sleeves = 20
	// Paper Straws
	steel.PaperStraws.Quantity = 2000
	steel.PaperStraws.Package = 250
	steel.PaperStraws.Packages = 8
	// Sip Lids
	steel.SipLids.Quantity = 1000
	steel.SipLids.Sleeve = 100
	steel.SipLids.Sleeves = 10
	// Ice Cream Lids
	steel.IceCreamLids.Quantity = 1000
	steel.IceCreamLids.Sleeve = 100
	steel.IceCreamLids.Sleeves = 10
	// Pepsi Lids
	steel.PepsiLids.Quantity = 2000
	steel.PepsiLids.Sleeve = 100
	steel.PepsiLids.Sleeves = 20
	// Chicken Bake Bags
	steel.ChickenBakeBags.Quantity = 1000
	// Pizza Boxes
	steel.PizzaBoxes.Quantity = 50

	var dailyPizzaBoxesTotal int = 0
	var dailyFeSmoothieTotal int = 0
	var dailyKioskSmoothieTotal int = 0
	var dailyFeColdBrewTotal int = 0
	var dailyKioskColdBrewTotal int = 0
	var dailyFeColdBrewLatteTotal int = 0
	var dailyKioskColdBrewLatteTotal int = 0
	var dailyFeSodaTotal int = 0
	var dailyKioskSodaTotal int = 0
	var dailyIceCreamTotal int = 0
	var dailyHotDogTotal int = 0

	// Get the daily total of items sold for the number of shop days.
	for i := 1; i < 5; i++ {
		// First get the id name in the html. Then the value and convert to integer and add to the daily total.
		// Convert string to number for each form value from index.html
		var pizzaSoldString string = "pizza_boxes" + strconv.Itoa(i)
		pizzaSold := r.FormValue(pizzaSoldString)
		pizzaBoxesTotal, _ := strconv.Atoi(pizzaSold)
		dailyPizzaBoxesTotal += pizzaBoxesTotal

		var answer string = "fe_smoothie" + strconv.Itoa(i)
		feSmoothieStringTotal := r.FormValue(answer)
		feSmoothieTotal, _ := strconv.Atoi(feSmoothieStringTotal)
		dailyFeSmoothieTotal += feSmoothieTotal

		var smoothieString string = "kiosk_smoothie" + strconv.Itoa(i)
		kioskSmoothieStringTotal := r.FormValue(smoothieString)
		kioskSmoothieTotal, _ := strconv.Atoi(kioskSmoothieStringTotal)
		dailyKioskSmoothieTotal += kioskSmoothieTotal

		var coldbrewString string = "fe_coldbrew" + strconv.Itoa(i)
		feColdBrewStringTotal := r.FormValue(coldbrewString)
		feColdBrewTotal, _ := strconv.Atoi(feColdBrewStringTotal)
		dailyFeColdBrewTotal += feColdBrewTotal

		coldbrewString = "kiosk_coldbrew" + strconv.Itoa(i)
		kioskColdBrewStringTotal := r.FormValue(coldbrewString)
		kioskColdBrewTotal, _ := strconv.Atoi(kioskColdBrewStringTotal)
		dailyKioskColdBrewTotal += kioskColdBrewTotal

		coldbrewString = "fe_coldbrew_latte" + strconv.Itoa(i)
		feColdBrewLatteStringTotal := r.FormValue(coldbrewString)
		feColdBrewLatteTotal, _ := strconv.Atoi(feColdBrewLatteStringTotal)
		dailyFeColdBrewLatteTotal += feColdBrewLatteTotal

		coldbrewString = "kiosk_coldbrew_latte" + strconv.Itoa(i)
		kioskColdBrewLatteStringTotal := r.FormValue(coldbrewString)
		kioskColdBrewLatteTotal, _ := strconv.Atoi(kioskColdBrewLatteStringTotal)
		dailyKioskColdBrewLatteTotal += kioskColdBrewLatteTotal

		var sodaString string = "fe_soda" + strconv.Itoa(i)
		feSodaStringTotal := r.FormValue(sodaString)
		feSodaTotal, _ := strconv.Atoi(feSodaStringTotal)
		dailyFeSodaTotal += feSodaTotal

		sodaString = "kiosk_soda" + strconv.Itoa(i)
		kioskSodaStringTotal := r.FormValue(sodaString)
		kioskSodaTotal, _ := strconv.Atoi(kioskSodaStringTotal)
		dailyKioskSodaTotal += kioskSodaTotal

		var iceCreamString string = "ice_cream_sales" + strconv.Itoa(i)
		iceCreamStringTotal := r.FormValue(iceCreamString)
		iceCreamTotal, _ := strconv.Atoi(iceCreamStringTotal)
		dailyIceCreamTotal += iceCreamTotal

		var hotdogString string = "hotdog_sales" + strconv.Itoa(i)
		hotdogStringTotal := r.FormValue(hotdogString)
		hotdogTotal, _ := strconv.Atoi(hotdogStringTotal)
		dailyHotDogTotal += hotdogTotal

	}

	/*
		kioskSmoothieString := r.FormValue("kiosk_smoothie")
		kioskSmoothie, _ := strconv.Atoi(kioskSmoothieString)

		feColdBrewString := r.FormValue("fe_coldbrew")
		feColdBrew, _ := strconv.Atoi(feColdBrewString)
		kioskColdBrewString := r.FormValue("kiosk_coldbrew")
		kioskColdBrew, _ := strconv.Atoi(kioskColdBrewString)
		feColdBrewLatteString := r.FormValue("fe_coldbrew_latte")
		feColdBrewLatte, _ := strconv.Atoi(feColdBrewLatteString)
		kioskColdBrewLatteString := r.FormValue("kiosk_coldbrew_latte")
		kioskColdBrewLatte, _ := strconv.Atoi(kioskColdBrewLatteString)
	*/
	var projectedPizzaBoxTotal int = int((.20 * float64(dailyPizzaBoxesTotal))) + dailyPizzaBoxesTotal
	pizzaBoxesNeeded := getBoxes(steel, projectedPizzaBoxTotal)
	// Get the amount of smoothie and cold brew sales
	var smoothieSales int = dailyKioskSmoothieTotal + dailyFeSmoothieTotal
	fmt.Printf("This is the total number of smoothies", smoothieSales)
	smoothieSales = int((.20 * float64(smoothieSales))) + smoothieSales
	var coldbrewSales int = dailyFeColdBrewTotal + dailyKioskColdBrewTotal + dailyFeColdBrewLatteTotal + dailyKioskColdBrewLatteTotal
	coldbrewSales = int((.20 * float64(coldbrewSales))) + coldbrewSales
	var sipLidTotal int = smoothieSales + coldbrewSales

	// Divide the total number of sip lids to get the sleeves we want.
	// Mod to get the remaider. Ex. 101 % 100 = 1 so we need two sleeves
	var sipLidsNeeded int = sipLidTotal / steel.SipLids.Sleeve
	remainder := sipLidTotal % steel.SipLids.Sleeve
	if remainder > 0 {
		sipLidsNeeded++
	}

	var hotdogSales int = int((.20 * float64(dailyHotDogTotal))) + dailyHotDogTotal

	var sodaLidsTotal int = dailyFeSodaTotal + dailyKioskSodaTotal
	sodaLidsTotal = int((.20 * float64(sodaLidsTotal))) + sodaLidsTotal + hotdogSales
	fmt.Printf("Soda lids", sodaLidsTotal)
	var sodaCupsTotal int = sodaLidsTotal

	var sodaLidsNeeded int = sodaLidsTotal / steel.PepsiLids.Sleeve
	var sodaCupsNeeded int = sodaCupsTotal / steel.SodaCups.Sleeve

	if sodaLidsTotal%steel.PepsiLids.Sleeve > 0 {
		sodaLidsNeeded++
	}

	if sodaCupsTotal%steel.SodaCups.Sleeve > 0 {
		sodaCupsNeeded++
	}

	var iceCreamSales int = int((.20 * float64(dailyIceCreamTotal))) + dailyIceCreamTotal

	var iceCreamLidsTotal int = iceCreamSales
	var sixteenOunceCupsTotal int = iceCreamLidsTotal + sipLidTotal
	var iceCreamLidsNeeded int = iceCreamLidsTotal / steel.IceCreamLids.Sleeve
	var sixteenOunceCupsNeeded int = sixteenOunceCupsTotal / steel.One6OzCups.Sleeve

	var paperStrawsTotal int = sodaCupsTotal + sipLidTotal
	var paperStrawsNeeded int = paperStrawsTotal / steel.PaperStraws.Package

	if paperStrawsTotal%steel.PaperStraws.Package > 0 {
		paperStrawsNeeded++
	}

	if iceCreamLidsTotal%steel.IceCreamLids.Sleeve > 0 {
		iceCreamLidsNeeded++
	}

	if sixteenOunceCupsTotal%steel.One6OzCups.Sleeve > 0 {
		sixteenOunceCupsNeeded++
	}

	// Add values to the struct
	data := ShopDetails{
		PizzaSold:              pizzaBoxesNeeded,
		FeSmoothieSold:         dailyFeSmoothieTotal,
		KioskSmoothieSold:      dailyKioskSmoothieTotal,
		FeColdBrewSold:         dailyFeColdBrewTotal,
		KioskColdBrewSold:      dailyKioskColdBrewTotal,
		FeColdBrewLatteSold:    dailyFeColdBrewLatteTotal,
		KioskColdBrewLatteSold: dailyKioskColdBrewLatteTotal,
		FeSodaSold:             dailyFeSodaTotal,
		KioskSodaSold:          dailyKioskSodaTotal,
		IceCreamSold:           dailyIceCreamTotal,
		HotDogSold:             dailyHotDogTotal,
		SipLidsToBuy:           sipLidsNeeded,
		SodaLidsToBuy:          sodaLidsNeeded,
		SodaCupsToBuy:          sodaCupsNeeded,
		IceCreamLidsToBuy:      iceCreamLidsNeeded,
		SixteenOunceCupsToBuy:  sixteenOunceCupsNeeded,
		PaperStrawsToBuy:       paperStrawsNeeded,
	}

	render(w, "templates/home.html", data)
}
