package pkg

import (
	"fmt"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/goqu_errors"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/utils"
)

func UpdateRowWithCondition(t entities.Table, conditionColumn entities.Column) (string, error) {
	baseQuery := "UPDATE %s SET %s WHERE %s = $%d"
	if !utils.ContainColumn(t, conditionColumn) {
		return "", &goqu_errors.TableDoesntContainColumn{T: t.GetName(), C: conditionColumn.Name}
	}
	clmns := t.GetColumns()
	utils.SortColumns(clmns)
	argsStr, lastPlaceHolderIndex := utils.MakeArgsStrForUpdate(clmns)
	return fmt.Sprintf(baseQuery, t.GetName(), argsStr, conditionColumn.Name, lastPlaceHolderIndex+1), nil
}
