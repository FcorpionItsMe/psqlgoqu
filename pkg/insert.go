package pkg

import (
	"fmt"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/utils"
)

func Insert(t entities.Table, notInsertableColumn entities.Column) string {
	baseQuery := "INSERT INTO %s (%s) VALUES (%s)"
	argsStr := utils.MakeParamPlaceholders(t.GetColumnsCount())
	clmns := t.GetColumns()
	utils.SortColumns(clmns)
	paramsStr := utils.MakeArgsStr(clmns)
	return fmt.Sprintf(baseQuery, t.GetName(), paramsStr, argsStr)
}
