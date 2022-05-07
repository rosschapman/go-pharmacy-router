package insurancePharmacy

import (
	"dispense/drug"
	"fmt"
	"math/rand"
	"strings"
)

type PharmacyInsurance struct {
	name      string
	inventory map[drug.DrugId]int
}

func (phc *PharmacyInsurance) Init(name string) {
	phc.name = name

	(phc.inventory) = make(map[drug.DrugId]int)

	for i := 0; i < len(drug.DrugMap); i++ {
		phc.LoadInventory((drug.DrugId(i)), rand.Intn(100))
	}
}

func (phc PharmacyInsurance) CheckInventory(drug drug.DrugId) (int, error) {
	fmt.Printf("Pharmacy(%s): Checking Inventory for %s\n", phc.name, drug)

	if quantity, ok := phc.inventory[drug]; ok {
		return quantity, nil
	} else {
		return -1, fmt.Errorf("%s not found in inventory", drug)
	}
}

func (phc PharmacyInsurance) DispenseDrug(drug drug.DrugId) error {
	fmt.Printf("Pharmacy(%s): Dispensing %s\n", phc.name, drug)
	quantity, err := phc.CheckInventory(drug)

	if err != nil {
		return err
	}

	if quantity > 0 {
		phc.inventory[drug] -= 1
	}

	return nil
}

func (phc PharmacyInsurance) LoadInventory(drug drug.DrugId, quantity int) {
	fmt.Printf("Pharmacy(%s): Loading Inventory Item %s\n", phc.name, drug)

	_, err := phc.CheckInventory(drug)

	if err != nil {
		phc.inventory[drug] = 0
	}

	phc.inventory[drug] += quantity
}

func (phc PharmacyInsurance) DisplayFullInventory() {
	var inventory strings.Builder
	for drug, quantity := range phc.inventory {
		fmt.Fprintf(&inventory, "\t%s Quantity: %d\n", drug, quantity)
	}

	fmt.Printf("Pharmacy: %s\nInventory [\n%s\n]\n", phc.name, &inventory)
}
