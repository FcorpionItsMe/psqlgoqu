package generator

import (
	"fmt"
	"strconv"
	"strings"
)

var mainTableTemplate = `
var %%TableVarName%% = %%TableName%%Table{
     Name: "%%TableName%%",
     Columns: %%TableName%%TableColumns{
               %%ColumnsInitialization%%
     },
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
	withStructColumns := strings.ReplaceAll(withColumnsCount, "%%Columns%%", GenerateColumnsForStruct(columns))
	withGetColumnsFuncColumns := strings.ReplaceAll(withStructColumns, "%%ColumnsForGetColumn%%", GenerateColumnsForGetColumns(columns))
	withInitialization := strings.ReplaceAll(withGetColumnsFuncColumns, "%%ColumnsInitialization%%", GenerateColumnsInitialization(columns))
	return withInitialization
}
func GenerateColumnsForStruct(columnsToGen []string) string {
	return strings.Join(columnsToGen, " entities.Column\n") + " entities.Column\n"
}

// Just comment!
func GenerateColumnsForGetColumns(columnsToGen []string) string {
	var res []string
	for i := 0; i < len(columnsToGen); i++ {
		res = append(res, "t."+columnsToGen[i])
	}
	return strings.Join(res, ", ")
}
func GenerateColumnsInitialization(clmns []string) string {
	var res []string
	for i := 0; i < len(clmns); i++ {
		column := clmns[i] + ": " + fmt.Sprintf("entities.Column{Name: pq.QuoteIdentifier(\"%s\"), Position: %d}", clmns[i], i)
		res = append(res, column)
	}
	return strings.Join(res, ",\n") + ",\n"
}
