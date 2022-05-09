package main

import (
	"dispense/drugs"
	"dispense/pharmacy"
	"dispense/router"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	drugs.LoadDrugMap(10)

	phCash := &pharmacy.PharmacyCash{}
	phCash.Init("AdamGreens", 10)

	phIns := &pharmacy.PharmacyInsurance{}
	phIns.Init("Rite Ross", 10)

	phCash.DisplayFullInventory()
	phIns.DisplayFullInventory()

	pRouter := router.PharmacyRouter{Cash: phCash, Insurance: phIns}

	for i := 0; i < 11; i++ {
		err := pRouter.ProcessOrder(drugs.DrugId(i))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
	}
}
