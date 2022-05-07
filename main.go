package main

import (
	cashPharmacy "dispense/cash-pharmacy"
	"dispense/drug"
	"fmt"
)

type PharmacyBase interface {
	DispenseDrug(drug drug.DrugId) error
	LoadInventory(drug drug.DrugId, quantity int)
	CheckInventory(drug drug.DrugId) (int, error)
	DisplayFullInventory()
}

func main() {
	phc1 := cashPharmacy.PharmacyCash{Inventory: map[drug.DrugId]int{1: 10, 2: 20, 3: 30}}

	fmt.Println(phc1.CheckInventory(1))
	fmt.Println(phc1.DispenseDrug(4))
	fmt.Println(phc1.CheckInventory(1))
	fmt.Println(phc1.DispenseDrug(1))
	phc1.LoadInventory(2, 10)
	fmt.Println(phc1.CheckInventory(2))
	phc1.DisplayFullInventory()
}
