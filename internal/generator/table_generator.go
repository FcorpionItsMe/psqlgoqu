package generator

import (
	"strconv"
	"strings"
)

var mainTableTemplate = `
var %%TableVarName%% = %%TableName%%Table{

}

type %%TableName%%Table struct {
   Name string
   Columns %%TableName%%TableColumns
}
type %%TableName%%TableColumns struct { 
   %%Columns%%
}

    func (t %%TableName%%Table) GetName() string {
        return t.Name
    }
	func (t %%TableName%%Table) GetColumns() []Column {
         return []Column{%%ColumnsForGetColumn%%}
    }
	func (t %%TableName%%Table) GetColumnsCount() int {
      return %%ColumnsCount%%
    }
	func (t %%TableName%%Table) String() string {
      return t.Name
    }
`

func GenerateTableCode(tableName string, columns ...string) string {
	withTableName := strings.ReplaceAll(mainTableTemplate, "%%TableName%%", tableName)
	var tableNameWithFirstLetterUpper string
	tableNameWithFirstLetterUpper += strings.ToUpper(string(tableName[0]))
	for i := 1; i < len(tableName); i++ {
		tableNameWithFirstLetterUpper += string(tableName[i])
	}
	withVarName := strings.ReplaceAll(withTableName, "%%TableVarName%%", tableNameWithFirstLetterUpper)
	columnsCount := len(columns)
	withColumnsCount := strings.ReplaceAll(withVarName, "%%ColumnsCount%%", strconv.Itoa(columnsCount))
	return withColumnsCount
}
