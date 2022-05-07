package drug

import (
	"fmt"
)

type DrugId int

var drugMap = map[DrugId]string{1: "drugA", 2: "drugB", 3: "drugC"}

func (d DrugId) String() string {
	return fmt.Sprintf("Drug #%d (%s)", d, drugMap[d])
}
