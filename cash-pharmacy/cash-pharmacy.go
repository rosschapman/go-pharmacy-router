package cashPharmacy

import (
	"dispense/drug"
	"fmt"
)

type PharmacyCash struct {
	Inventory map[drug.DrugId]int
}

func (phc PharmacyCash) CheckInventory(drug drug.DrugId) (int, error) {
	fmt.Printf("Check Inventory for %s\n", drug)

	if quantity, ok := phc.Inventory[drug]; ok {
		return quantity, nil
	} else {
		return -1, fmt.Errorf("%s not found in inventory", drug)
	}
}

func (phc PharmacyCash) DispenseDrug(drug drug.DrugId) error {
	fmt.Printf("Dispense drug for %s\n", drug)
	amount, err := phc.CheckInventory(drug)

	if err != nil {
		return err
	}

	if amount > 0 {
		phc.Inventory[drug] -= 1
	}

	return nil
}

func (phc PharmacyCash) LoadInventory(drug drug.DrugId, quantity int) {
	_, err := phc.CheckInventory(drug)

	if err != nil {
		phc.Inventory[drug] = 0
	}

	phc.Inventory[drug] += quantity
}

func (phc PharmacyCash) DisplayFullInventory() {
	for drug, quantity := range phc.Inventory {
		fmt.Printf("%s Quantity: %d\n", drug, quantity)
	}
}
