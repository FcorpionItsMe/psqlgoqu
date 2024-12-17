package entities

type Table interface {
	GetName() string
	GetColumns() []Column
	GetColumnsCount() int
	String() string
}
