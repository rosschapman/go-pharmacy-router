package drug

import (
	"fmt"
)

type DrugId int

var DrugMap = map[DrugId]string{0: "drugA", 1: "drugB", 2: "drugC", 3: "drugD"}

func (drugId DrugId) String() string {
	if drugName, ok := DrugMap[drugId]; ok {
		return fmt.Sprintf("Drug #%d (%s)", drugId, drugName)
	}

	return fmt.Sprintf("Drug #%d (Drug Undefined)", drugId)
}
