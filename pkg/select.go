package pkg

import (
	"fmt"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/goqu_errors"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/utils"
	"strconv"
)

func SelectAll(t entities.Table) string {
	baseQuery := "SELECT * FROM %s"
	return fmt.Sprintf(baseQuery, t.GetName())
}
func SelectByColumnCondition(t entities.Table, conditionColumn entities.Column, val interface{}) (string, error) {
	if !utils.ContainColumn(t, conditionColumn) {
		return "", &goqu_errors.TableDoesntContainColumn{T: t.GetName(), C: conditionColumn.Name}
	}
	baseQuery := "SELECT * FROM %s WHERE %s = "
	if str, ok := val.(string); ok {
		baseQuery += str
	} else if num, ok := val.(int); ok {
		baseQuery += strconv.Itoa(num)
	}
	return fmt.Sprintf(baseQuery, t.GetName(), conditionColumn.Name), nil
}

///TODO: Сделать метод универсальнее сделав column -> slice Column ([]Column)
