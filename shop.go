package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Steel struct {
	SmoothieBase struct {
		Quantity     int `json:"quantity"`
		One6OzGallon int `json:"16oz/gallon"`
	} `json:"smoothie base"`
	DairyMix struct {
		Quantity     int `json:"quantity"`
		One6OzGallon int `json:"16oz/gallon"`
	} `json:"dairy mix"`
	SodaCups struct {
		Quantity int `json:"quantity"`
		Sleeve   int `json:"sleeve"`
		Sleeves  int `json:"sleeves"`
	} `json:"soda cups"`
	One6OzCups struct {
		Quantity int `json:"quantity"`
		Sleeve   int `json:"sleeve"`
		Sleeves  int `json:"sleeves"`
	} `json:"16 oz cups"`
	PaperStraws struct {
		Quantity int `json:"quantity"`
		Package  int `json:"package"`
		Packages int `json:"packages"`
	} `json:"paper straws"`
	SipLids struct {
		Quantity int `json:"quantity"`
		Sleeve   int `json:"sleeve"`
		Sleeves  int `json:"sleeves"`
	} `json:"sip lids"`
	IceCreamLids struct {
		Quantity int `json:"quantity"`
		Sleeve   int `json:"sleeve"`
		Sleeves  int `json:"sleeves"`
	} `json:"ice cream lids"`
	PepsiLids struct {
		Quantity int `json:"quantity"`
		Sleeve   int `json:"sleeve"`
		Sleeves  int `json:"sleeves"`
	} `json:"pepsi lids"`
	ChickenBakeBags struct {
		Quantity int `json:"quantity"`
	} `json:"chicken bake bags"`
	PizzaBoxes struct {
		Quantity int `json:"quantity"`
	} `json:"pizza boxes"`
}

/*
	Get the number of pizza sales for the days till next shop.
*/
func getPizzaSales(product Steel) int {
	//******************** Pizza Sales ********************//
	fmt.Println("How many pizzas sold until next shop day?")
	var pizzaSales int
	fmt.Scanln(&pizzaSales)

	totalPizzaBoxes := pizzaSales / product.PizzaBoxes.Quantity
	// Check to see if we need another set of boxes. If no remainder then we don't need another set.
	remainder := pizzaSales % product.PizzaBoxes.Quantity

	if remainder > 0 {
		totalPizzaBoxes++
	}

	return totalPizzaBoxes
}

/*
	Get the number beverages sold.
*/
func getBeverageSales(beverageSales map[string]int) (int, int, int) {

	for k := range beverageSales {
		fmt.Printf("How much of %s were sold? ", k)
		var bevSales int
		fmt.Scanln(&bevSales)
		beverageSales[k] = bevSales
	}

	smoothieSales := beverageSales["feFruitSmoothie"] + beverageSales["kioskFruitSmoothie"] + beverageSales["fcFruitSmoothie"]
	coldBrewSales := beverageSales["feColdBrew"] + beverageSales["kioskColdBrew"] + beverageSales["fcColdBrew"]
	sodaSales := beverageSales["feSoda"] + beverageSales["kioskSoda"] + beverageSales["fcSoda"]

	return smoothieSales, coldBrewSales, sodaSales
}

/*
	Get ice cream sales.
*/
func getIceCreamSales() int {
	var iceCreamSales int
	fmt.Println("How many ice creams sold? ")
	fmt.Scanln(&iceCreamSales)

	return iceCreamSales
}

/*
	Get the hot dogs sold.
*/
func getHotDogSales() int {
	var hotDogSales int
	fmt.Println("How many hot dogs sold? ")
	fmt.Scanln(&hotDogSales)

	return hotDogSales
}

/*
	Get the number of ice cream lids needed.
*/
func getIceCreamLidsNeeded(iceCreamSales int, product Steel) int {
	iceCremLidsNeeded := iceCreamSales / product.IceCreamLids.Sleeve
	if iceCremLidsNeeded%product.IceCreamLids.Sleeve > 0 {
		iceCremLidsNeeded++
	}

	return iceCremLidsNeeded
}

/*
	Get the number of sip lids needed for the smoothie and latte.
*/
func getSipLidsNeeded(smoothieSales int, coldBrewSales int, product Steel) int {
	sipLidsNeeded := (smoothieSales + coldBrewSales) / product.SipLids.Sleeve
	if sipLidsNeeded%product.SipLids.Sleeve > 0 {
		sipLidsNeeded++
	}

	return sipLidsNeeded
}

/*
	Get the number of 16 ounce cups needed.
*/
func getSixteenOunceCupsNeeded(smoothieSales int, coldBrewSales int, iceCreamSales int, product Steel) int {
	sixteenOunceCupsNeeded := (smoothieSales + coldBrewSales + iceCreamSales) / product.One6OzCups.Sleeve

	if sixteenOunceCupsNeeded%product.One6OzCups.Sleeve > 0 {
		sixteenOunceCupsNeeded++
	}

	return sixteenOunceCupsNeeded
}

/*
	Get the straws needed for the drinks.
*/
func getStrawsNeeded(drinksSold int, product Steel) int {
	strawsNeeded := drinksSold / product.PaperStraws.Package
	if strawsNeeded%product.PaperStraws.Package > 0 {
		strawsNeeded++
	}
	return strawsNeeded
}

func getSodaCupsNeeded(sodaSales int, hotDogSales int, product Steel) int {
	sodaCupsNeeded := (sodaSales + hotDogSales) / product.SodaCups.Sleeve
	if sodaCupsNeeded%product.SodaCups.Sleeve > 0 {
		sodaCupsNeeded++
	}

	return sodaCupsNeeded

}

func getBoxes(product Steel, pizzaSold int) int {
	totalPizzaBoxes := pizzaSold / product.PizzaBoxes.Quantity
	// Check to see if we need another set of boxes. If no remainder then we don't need another set.
	remainder := pizzaSold % product.PizzaBoxes.Quantity

	if remainder > 0 {
		totalPizzaBoxes++
	}

	fmt.Println(totalPizzaBoxes)
	return totalPizzaBoxes
}

// Render page. tmpl is the location of the file ex. templates/index.html
func render(w http.ResponseWriter, tmpl string, data ShopDetails) {
	// Parse the html and check if there is an error.
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Execute template and pass in the struct and check if there is an error
	err = t.Execute(w, data)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

// Render page. tmpl is the location of the file ex. templates/index.html
func renderInitialPage(w http.ResponseWriter, tmpl string, data ShopDays) {
	// Parse the html and check if there is an error.
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Execute template and pass in the struct and check if there is an error
	err = t.Execute(w, data)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

/******************* Main *********************/
func main() {

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
	steel.PaperStraws.Package = 500
	steel.PaperStraws.Packages = 4
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

	//fmt.Println(steel.SmoothieBase.Quantity)
	//fmt.Printf("%+v", steel)

	totalPizzaBoxes := getPizzaSales(steel)
	fmt.Printf("You will need %d pizza boxes\n", totalPizzaBoxes)

	//******************** Beverages ********************//
	fmt.Println("***** Checking all the beverages sold *****")
	//var beverageSales map[string]int
	beverageSales := make(map[string]int)
	beverageSales["feFruitSmoothie"] = 0
	beverageSales["kioskFruitSmoothie"] = 0
	beverageSales["fcFruitSmoothie"] = 0
	beverageSales["feColdBrew"] = 0
	beverageSales["kioskColdBrew"] = 0
	beverageSales["fcColdBrew"] = 0
	beverageSales["feSoda"] = 0
	beverageSales["kioskSoda"] = 0
	beverageSales["fcSoda"] = 0

	smoothieSales, coldBrewSales, sodaSales := getBeverageSales(beverageSales)
	// Smoothie and cold brew sales for a Sunday are 164 and individual sodas is 175
	fmt.Println(smoothieSales + coldBrewSales)
	fmt.Println(sodaSales)

	//******************** Smoothie/Latte/Ice Cream lids ********************//
	var iceCreamSales int = getIceCreamSales()

	//******************** Hot Dogs ********************//
	fmt.Println("***** Checking the hot dogs sold *****")
	var hotDogSales int = getHotDogSales()

	fmt.Println("\n")

	fmt.Println("***** Checking cups/lids needed for smoothie/latte/ice cream *****")
	iceCremLidsNeeded := getIceCreamLidsNeeded(iceCreamSales, steel)

	fmt.Printf("You going to need %d ice creeam sleeves.\n", iceCremLidsNeeded)

	sipLipsNeeded := getSipLidsNeeded(smoothieSales, coldBrewSales, steel)
	sixteenOunceCupsNeeded := getSixteenOunceCupsNeeded(smoothieSales, coldBrewSales, iceCreamSales, steel)

	fmt.Printf("You going to need %d sip lid sleeves. \n", sipLipsNeeded)
	fmt.Printf("You going to need %d 16 oz cups sleeves. \n", sixteenOunceCupsNeeded)

	//******************** Paper Straws ********************//
	fmt.Println("***** Checking the paper straws *****")
	drinksSold := smoothieSales + coldBrewSales + sodaSales + hotDogSales
	strawsNeeded := getStrawsNeeded(drinksSold, steel)

	totalStraws := (drinksSold * 2) / steel.PaperStraws.Quantity
	fmt.Printf("You going to need %d boxes of straws.\n", totalStraws)
	fmt.Printf("You going to need %d packages.\n", strawsNeeded)

	sodaCupsNeeded := getSodaCupsNeeded(sodaSales, hotDogSales, steel)
	fmt.Printf("You going to need %d soda sleeves.\n", sodaCupsNeeded)

	// Used to get css on the pages.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/home", HomePage)

	//log.Fatal(http.ListenAndServe(":8080", nil))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
