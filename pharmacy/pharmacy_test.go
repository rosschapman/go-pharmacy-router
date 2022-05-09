package pharmacy

import (
	"dispense/drugs"
	"fmt"
	"testing"
)

const GLOBAL_QUANTITY int = 10

func TestInsuranceInit(t *testing.T) {
	t.Run("initialize pharmacy", func(t *testing.T) {
		drugs.LoadDrugMap(2)
		pharmacy := PharmacyInsurance{}

		pharmacyName := "Pharmacy Insurance Name"
		pharmacy.Init(pharmacyName, GLOBAL_QUANTITY)

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

func TestInsuranceCheckInventory(t *testing.T) {
	drugs.LoadDrugMap(1)
	pharmacy := PharmacyInsurance{}
	pharmacy.Init("Pharmacy Insurance Name", GLOBAL_QUANTITY)

	t.Run("can find inventory item", func(t *testing.T) {
		result, _ := pharmacy.CheckInventory(0)
		expected := pharmacy.inventory[0]

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("cannot find inventory item", func(t *testing.T) {
		var drugId drugs.DrugId = 1
		_, result := pharmacy.CheckInventory(drugId)
		expected := fmt.Errorf("%s not found in inventory", drugId)

		if result == nil {
			t.Errorf("result %q expected %q", result, expected)
		}
	})
}

func TestInsuranceDispenseDrug(t *testing.T) {
	drugs.LoadDrugMap(1)
	pharmacy := PharmacyInsurance{}
	pharmacy.Init("Pharmacy Insurance Name", GLOBAL_QUANTITY)

	t.Run("dispenses a drug", func(t *testing.T) {
		var drugId drugs.DrugId = 0
		pharmacy.DispenseDrug(drugId)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := pharmacy.inventory[drugId]

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("does not dispense a drug and does not throw error", func(t *testing.T) {
		var drugId drugs.DrugId = 0
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
		var drugId drugs.DrugId = 1
		pharmacy.DispenseDrug(drugId)

		_, err := pharmacy.CheckInventory(drugId)
		if err == nil {
			t.Errorf("%s was found in inventory", drugId)
		}
	})
}

func TestInsuranceLoadInventory(t *testing.T) {
	drugs.LoadDrugMap(1)
	pharmacy := PharmacyInsurance{}
	pharmacy.Init("Pharmacy Insurance Name", GLOBAL_QUANTITY)

	t.Run("increases inventory stock", func(t *testing.T) {
		var drugId drugs.DrugId = 0
		pharmacy.LoadInventory(drugId, 1)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := pharmacy.inventory[drugId]

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("adds drug inventory", func(t *testing.T) {
		var drugId drugs.DrugId = 0
		delete(pharmacy.inventory, drugId)

		pharmacy.LoadInventory(drugId, 1)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := 1

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
}

func TestInsuranceDisplayFullInventory(t *testing.T) {
	drugs.LoadDrugMap(1)
	pharmacy := PharmacyInsurance{}
	pharmacy.Init("Pharmacy Insurance Name", GLOBAL_QUANTITY)

	t.Run("display full inventory", func(t *testing.T) {
		pharmacy.DisplayFullInventory()
	})
}

func TestCashInit(t *testing.T) {
	t.Run("initialize pharmacy", func(t *testing.T) {
		drugs.LoadDrugMap(2)
		pharmacy := PharmacyCash{}

		pharmacyName := "Pharmacy Cash Name"
		pharmacy.Init(pharmacyName, GLOBAL_QUANTITY)

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

func TestCashCheckInventory(t *testing.T) {
	drugs.LoadDrugMap(1)
	pharmacy := PharmacyCash{}
	pharmacy.Init("Pharmacy Cash Name", GLOBAL_QUANTITY)

	t.Run("can find inventory item", func(t *testing.T) {
		result, _ := pharmacy.CheckInventory(0)
		expected := pharmacy.inventory[0]

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("cannot find inventory item", func(t *testing.T) {
		var drugId drugs.DrugId = 1
		_, result := pharmacy.CheckInventory(drugId)
		expected := fmt.Errorf("%s not found in inventory", drugId)

		if result == nil {
			t.Errorf("result %q expected %q", result, expected)
		}
	})
}

func TestCashDispenseDrug(t *testing.T) {
	drugs.LoadDrugMap(1)
	pharmacy := PharmacyCash{}
	pharmacy.Init("Pharmacy Cash Name", GLOBAL_QUANTITY)

	t.Run("dispenses a drug", func(t *testing.T) {
		var drugId drugs.DrugId = 0
		pharmacy.DispenseDrug(drugId)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := pharmacy.inventory[drugId]

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("does not dispense a drug and does not throw error", func(t *testing.T) {
		var drugId drugs.DrugId = 0
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
		var drugId drugs.DrugId = 1
		pharmacy.DispenseDrug(drugId)

		_, err := pharmacy.CheckInventory(drugId)
		if err == nil {
			t.Errorf("%s was found in inventory", drugId)
		}
	})
}

func TestCashLoadInventory(t *testing.T) {
	drugs.LoadDrugMap(1)
	pharmacy := PharmacyCash{}
	pharmacy.Init("Pharmacy Cash Name", GLOBAL_QUANTITY)

	t.Run("increases inventory stock", func(t *testing.T) {
		var drugId drugs.DrugId = 0
		pharmacy.LoadInventory(drugId, 1)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := pharmacy.inventory[drugId]

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	t.Run("adds drug inventory", func(t *testing.T) {
		var drugId drugs.DrugId = 0
		delete(pharmacy.inventory, drugId)

		pharmacy.LoadInventory(drugId, 1)

		result, _ := pharmacy.CheckInventory(drugId)
		expected := 1

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
}

func TestCashDisplayFullInventory(t *testing.T) {
	drugs.LoadDrugMap(1)
	pharmacy := PharmacyCash{}
	pharmacy.Init("Pharmacy Cash Name", GLOBAL_QUANTITY)

	t.Run("display full inventory", func(t *testing.T) {
		pharmacy.DisplayFullInventory()
	})
}
