package main

import (
	"dispense/drug"
	cashPharmacy "dispense/pharmacy/cash"
	insurancePharmacy "dispense/pharmacy/insurance"
	"math/rand"
	"time"
)

type PharmacyBase interface {
	DispenseDrug(drug drug.DrugId) error
	LoadInventory(drug drug.DrugId, quantity int)
	CheckInventory(drug drug.DrugId) (int, error)
	DisplayFullInventory()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	phc1 := cashPharmacy.PharmacyCash{}
	phc1.Init("AdamsGreens")

	phc2 := insurancePharmacy.PharmacyInsurance{}
	phc2.Init("Rite Ross")

	phc1.DisplayFullInventory()
	phc2.DisplayFullInventory()
}
