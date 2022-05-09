package router

import (
	"dispense/drugs"
	"dispense/pharmacy"
	"fmt"
	"math/rand"
)

type PharmacyRouter struct {
	Cash      *pharmacy.PharmacyCash
	Insurance *pharmacy.PharmacyInsurance
}

func (prouter PharmacyRouter) ProcessOrder(drug drugs.DrugId) error {
	if _, ok := drugs.DrugMap[drug]; ok {
		cashQuantity, err := prouter.Cash.CheckInventory(drug)
		if err != nil {
			fmt.Println(err)
		}

		insuranceQuantity, err := prouter.Insurance.CheckInventory(drug)
		if err != nil {
			fmt.Println(err)
		}

		if cashQuantity < 1 && insuranceQuantity < 1 {
			return fmt.Errorf("no inventory to process order for %s", drug)
		}

		pharmacyToDispense := 0

		if cashQuantity > insuranceQuantity {
			pharmacyToDispense = 0
		} else if insuranceQuantity > cashQuantity {
			pharmacyToDispense = 1
		} else {
			cashOrInsurance := rand.Intn(2)

			if cashOrInsurance == 0 {
				pharmacyToDispense = 0
			} else {
				pharmacyToDispense = 1
			}
		}

		if pharmacyToDispense == 0 {
			err := assignOrder(prouter.Cash, drug)
			if err != nil {
				return err
			}
		} else {
			err := assignOrder(prouter.Insurance, drug)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return fmt.Errorf("could not find %s in DrugMap", drug)
}

func assignOrder(p pharmacy.Pharmacy, drug drugs.DrugId) error {
	err := p.DispenseDrug(drug)
	if err != nil {
		return err
	}

	return nil
}
