package drugs

import (
	"fmt"
	"testing"
)

func TestLoadDrugMap(t *testing.T) {
	t.Run("loads drug map with 10 items", func(t *testing.T) {
		LoadDrugMap(2)
		result := len(DrugMap)
		expected := 2

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
}

func TestString(t *testing.T) {
	LoadDrugMap(1)

	t.Run("drug name found", func(t *testing.T) {
		var drugId DrugId = 0
		result := fmt.Sprint(drugId)
		expected := fmt.Sprintf("DrugRx_%d", drugId)

		if result != expected {
			t.Errorf("result %s expected %s", result, expected)
		}
	})

	t.Run("drug name not found", func(t *testing.T) {
		var drugId DrugId = 1
		result := fmt.Sprint(drugId)
		expected := fmt.Sprintf("No Drug Found for #%d", drugId)

		if result != expected {
			t.Errorf("result %s expected %s", result, expected)
		}
	})
}
