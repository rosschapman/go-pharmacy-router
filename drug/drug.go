package drug

import (
	"fmt"
)

type DrugId int

var DrugMap = map[DrugId]string{}

func LoadDrugMap(drugMapSize int) {
	DrugMap = map[DrugId]string{}

	for i := 0; i < drugMapSize; i++ {
		DrugMap[DrugId(i)] = fmt.Sprintf("DrugRx_%d", i)
	}
}

func (drugId DrugId) String() string {
	if drugName, ok := DrugMap[drugId]; ok {
		return drugName
	}

	return fmt.Sprintf("No Drug Found for #%d", drugId)
}
