package pkg

import (
	"fmt"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/goqu_errors"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/utils"
)

func SelectAll(t entities.Table) string {
	baseQuery := "SELECT * FROM %s"
	return fmt.Sprintf(baseQuery, t.GetName())
}
func SelectByColumnCondition(t entities.Table, conditionColumn entities.Column) (string, error) {
	if !utils.ContainColumn(t, conditionColumn) {
		return "", &goqu_errors.TableDoesntContainColumn{T: t.GetName(), C: conditionColumn.Name}
	}
	baseQuery := "SELECT * FROM %s WHERE %s = $1"
	return fmt.Sprintf(baseQuery, t.GetName(), conditionColumn.Name), nil
}

///TODO: Сделать метод универсальнее сделав column -> slice Column ([]Column)
