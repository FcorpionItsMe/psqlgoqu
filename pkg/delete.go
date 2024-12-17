package pkg

import (
	"fmt"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/goqu_errors"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/utils"
)

func DeleteRowWithCondition(t entities.Table, conditionColumn entities.Column) (string, error) {
	if !utils.ContainColumn(t, conditionColumn) {
		return "", &goqu_errors.TableDoesntContainColumn{T: t.GetName(), C: conditionColumn.Name}
	}
	baseQuery := "DELETE FROM %s WHERE %s = $1"
	return fmt.Sprintf(baseQuery, t.GetName(), conditionColumn), nil
}
