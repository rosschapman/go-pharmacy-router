package pharmacy

import (
	"dispense/drugs"
	"fmt"
	"math/rand"
	"strings"
)

type Pharmacy interface {
	Init(name string, quantity int)
	DispenseDrug(drug drugs.DrugId) error
	LoadInventory(drug drugs.DrugId, quantity int)
	CheckInventory(drug drugs.DrugId) (int, error)
	DisplayFullInventory()
}

type PharmacyCash struct {
	name      string
	inventory map[drugs.DrugId]int
}

func (phc *PharmacyCash) Init(name string, quantity int) {
	phc.name = name

	phc.inventory = make(map[drugs.DrugId]int)

	for i := 0; i < len(drugs.DrugMap); i++ {
		phc.LoadInventory((drugs.DrugId(i)), rand.Intn(quantity))
	}
}

func (phc PharmacyCash) CheckInventory(drug drugs.DrugId) (int, error) {
	fmt.Printf("Pharmacy(%s): Checking Inventory for %s\n", phc.name, drug)

	if quantity, ok := phc.inventory[drug]; ok {
		return quantity, nil
	} else {
		return -1, fmt.Errorf("%s not found in inventory", drug)
	}
}

func (phc PharmacyCash) DispenseDrug(drug drugs.DrugId) error {
	quantity, err := phc.CheckInventory(drug)

	if err != nil {
		return err
	}

	if quantity > 0 {
		phc.inventory[drug] -= 1
		fmt.Printf("Pharmacy(%s): Dispensing %s with Cash Payment\n", phc.name, drug)
	} else {
		fmt.Printf("Pharmacy(%s): No more inventory for %s\n", phc.name, drug)
	}

	return nil
}

func (phc PharmacyCash) LoadInventory(drug drugs.DrugId, quantity int) {
	_, err := phc.CheckInventory(drug)
	fmt.Printf("Pharmacy(%s): Loading Inventory Item %s\n", phc.name, drug)

	if err != nil {
		phc.inventory[drug] = 0
	}

	phc.inventory[drug] += quantity
}

func (phc PharmacyCash) DisplayFullInventory() {
	var inventory strings.Builder
	for drug, quantity := range phc.inventory {
		fmt.Fprintf(&inventory, "    %s - Quantity: %d\n", drug, quantity)
	}

	fmt.Printf("Pharmacy: {\n  name: %s\n  inventory: [\n%s\n  ]\n}\n", phc.name, &inventory)
}

type PharmacyInsurance struct {
	name      string
	inventory map[drugs.DrugId]int
}

func (phc *PharmacyInsurance) Init(name string, quantity int) {
	phc.name = name

	phc.inventory = make(map[drugs.DrugId]int)

	for i := 0; i < len(drugs.DrugMap); i++ {
		phc.LoadInventory((drugs.DrugId(i)), rand.Intn(quantity))
	}
}

func (phc PharmacyInsurance) CheckInventory(drug drugs.DrugId) (int, error) {
	fmt.Printf("Pharmacy(%s): Checking Inventory for %s\n", phc.name, drug)

	if quantity, ok := phc.inventory[drug]; ok {
		return quantity, nil
	} else {
		return -1, fmt.Errorf("%s not found in inventory", drug)
	}
}

func (phc PharmacyInsurance) DispenseDrug(drug drugs.DrugId) error {
	quantity, err := phc.CheckInventory(drug)

	if err != nil {
		return err
	}

	if quantity > 0 {
		phc.inventory[drug] -= 1
		fmt.Printf("Pharmacy(%s): Dispensing %s with Insurance Payment\n", phc.name, drug)
	} else {
		fmt.Printf("Pharmacy(%s): No more inventory for %s\n", phc.name, drug)
	}

	return nil
}

func (phc PharmacyInsurance) LoadInventory(drug drugs.DrugId, quantity int) {
	_, err := phc.CheckInventory(drug)
	fmt.Printf("Pharmacy(%s): Loading Inventory Item %s\n", phc.name, drug)

	if err != nil {
		phc.inventory[drug] = 0
	}

	phc.inventory[drug] += quantity
}

func (phc PharmacyInsurance) DisplayFullInventory() {
	var inventory strings.Builder
	for drug, quantity := range phc.inventory {
		fmt.Fprintf(&inventory, "    %s - Quantity: %d\n", drug, quantity)
	}

	fmt.Printf("Pharmacy: {\n  name: %s\n  inventory: [\n%s\n  ]\n}\n", phc.name, &inventory)
}
