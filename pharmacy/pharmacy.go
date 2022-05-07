package pharmacy

import "dispense/drug"

type Pharmacy interface {
	Init()
	DispenseDrug(drug drug.DrugId) error
	LoadInventory(drug drug.DrugId, quantity int)
	CheckInventory(drug drug.DrugId) (int, error)
	DisplayFullInventory()
}
