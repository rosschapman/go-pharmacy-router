package router

import (
	"dispense/drugs"
	"dispense/pharmacy"
	"math/rand"
	"testing"
	"time"
)

const GLOBAL_QUANTITY int = 10

func TestProcessOrder(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	drugs.LoadDrugMap(1)
	phCash := &pharmacy.PharmacyCash{}
	phCash.Init("AdamGreens", GLOBAL_QUANTITY)
	phIns := &pharmacy.PharmacyInsurance{}
	phIns.Init("Rite Ross", GLOBAL_QUANTITY)
	pRouter := PharmacyRouter{Cash: phCash, Insurance: phIns}

	t.Run("drug is not in cash inventory", func(t *testing.T) {
		err := pRouter.ProcessOrder(1)
		if err == nil {
			t.Errorf("drug does not exist, should error out")
		}
	})

	t.Run("drug is not in insurance inventory", func(t *testing.T) {
		err := pRouter.ProcessOrder(1)
		if err == nil {
			t.Errorf("drug does not exist, should error out")
		}
	})

	t.Run("drug quantity too low or is not in either inventory", func(t *testing.T) {

	})

	t.Run("cash has more inventory than insurance", func(t *testing.T) {

	})

	t.Run("insurance has more inventory than cash", func(t *testing.T) {

	})

	t.Run("insurance has more inventory than cash", func(t *testing.T) {

	})

	t.Run("insurance has same inventory as cash, assign to cash", func(t *testing.T) {

	})

	t.Run("insurance has same inventory as cash, assign to insurance", func(t *testing.T) {

	})

	t.Run("cash inventory insufficient, cannot dispense", func(t *testing.T) {

	})

	t.Run("insurance inventory insufficient, cannot dispense", func(t *testing.T) {

	})

	t.Run("drug is not in inventory", func(t *testing.T) {
		err := pRouter.ProcessOrder(1)
		if err == nil {
			t.Errorf("drug does not exist, should error out")
		}
	})
}

func TestAssignOrder(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	drugs.LoadDrugMap(1)
	phCash := &pharmacy.PharmacyCash{}
	phCash.Init("AdamGreens", GLOBAL_QUANTITY)
	phIns := &pharmacy.PharmacyInsurance{}
	phIns.Init("Rite Ross", GLOBAL_QUANTITY)
	pRouter := PharmacyRouter{Cash: phCash, Insurance: phIns}

	t.Run("dispenses drug", func(t *testing.T) {
		err := assignOrder(pRouter.Cash, 0)
		if err != nil {
			t.Errorf("drug should have been dispensed")
		}
	})

	t.Run("does not dispense drug", func(t *testing.T) {
		err := assignOrder(pRouter.Cash, 1)
		if err == nil {
			t.Errorf("drug should not have been dispensed")
		}
	})
}
