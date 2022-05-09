package insurancePharmacy

import (
	"dispense/drug"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("initialize pharmacy", func(t *testing.T) {
		drug.LoadDrugMap(2)
		pharmacy := PharmacyInsurance{}

		pharmacyName := "Pharmacy Name"
		pharmacy.Init(pharmacyName)

		result := len(pharmacy.inventory)
		expected := 2

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}

		resultName := pharmacy.name

		if resultName != pharmacyName {
			t.Errorf("result %s expected %s", resultName, pharmacyName)
		}
	})
}

func TestCheckInventory(t *testing.T) {
	drug.LoadDrugMap(1)
	pharmacy := PharmacyInsurance{}
	pharmacy.Init("pharmacy name")

	t.Run("can find inventory item", func(t *testing.T) {
		result, _ := pharmacy.CheckInventory(0)
		expected := 10

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("can find inventory item", func(t *testing.T) {
		var drugId drug.DrugId = 1
		_, result := pharmacy.CheckInventory(drugId)
		expected := fmt.Errorf("%s not found in inventory", drugId)

		if result == nil {
			t.Errorf("result %q expected %q", result, expected)
		}
	})
}

func TestDispenseDrug(t *testing.T) {
	drug.LoadDrugMap(1)
	pharmacy := PharmacyInsurance{}
	pharmacy.Init("pharmacy name")

	t.Run("dispenses a drug", func(t *testing.T) {
		var drugId drug.DrugId = 0
		pharmacy.DispenseDrug(drugId)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := 9

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("does not dispense a drug and does not throw error", func(t *testing.T) {
		var drugId drug.DrugId = 0
		pharmacy.inventory[drugId] = 0

		pharmacy.DispenseDrug(drugId)

		result, err := pharmacy.CheckInventory(drugId)
		if err != nil {
			t.Errorf("could not find %s in inventory", drugId)
		}

		if result != 0 {
			t.Errorf("%s dispensed although empty inventory", drugId)
		}
	})

	t.Run("does not dispense a drug and does not throw error", func(t *testing.T) {
		var drugId drug.DrugId = 1
		pharmacy.DispenseDrug(drugId)

		_, err := pharmacy.CheckInventory(drugId)
		if err == nil {
			t.Errorf("%s was found in inventory", drugId)
		}
	})
}

func TestLoadInventory(t *testing.T) {
	drug.LoadDrugMap(1)
	pharmacy := PharmacyInsurance{}
	pharmacy.Init("pharmacy name")

	t.Run("increases inventory stock", func(t *testing.T) {
		var drugId drug.DrugId = 0
		pharmacy.LoadInventory(drugId, 1)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := 11

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("adds drug inventory", func(t *testing.T) {
		var drugId drug.DrugId = 0
		delete(pharmacy.inventory, drugId)

		pharmacy.LoadInventory(drugId, 1)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := 1

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
}
