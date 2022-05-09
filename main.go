package main

import (
	"dispense/drug"
	cashPharmacy "dispense/pharmacy/cash-pharmacy"
	insurancePharmacy "dispense/pharmacy/insurance-pharmacy"
)

type PharmacyBase interface {
	DispenseDrug(drug drug.DrugId) error
	LoadInventory(drug drug.DrugId, quantity int)
	CheckInventory(drug drug.DrugId) (int, error)
	DisplayFullInventory()
}

func main() {
	drug.LoadDrugMap(10)

	phCash := cashPharmacy.PharmacyCash{}
	phCash.Init("AdamsGreens")

	phIns := insurancePharmacy.PharmacyInsurance{}
	phIns.Init("Rite Ross")

	phCash.DisplayFullInventory()
	phIns.DisplayFullInventory()

	phCash.DispenseDrug(1)
	phIns.DispenseDrug(3)

	phCash.DisplayFullInventory()
	phIns.DisplayFullInventory()
}
