package main

import (
	"fmt"
	"github.com/FcorpionItsMe/psqlgoqu/cmd"
	"github.com/FcorpionItsMe/psqlgoqu/internal/generator"
)

func main() {
	//repository, err := json_repository.NewRepository()
	//if err != nil {
	//	log.Fatal(err)
	//}
	////for i := 0; i < 10; i++ {
	////	err = repository.SavePath(domain.PathItem{
	////		ActualPath:    fmt.Sprintf("ActualPath %d", i),
	////		FirstCallPath: fmt.Sprintf("FirstPath %d", i),
	////		LastCallPath:  fmt.Sprintf("LastPath %d", i),
	////	})
	////	if err != nil {
	////		log.Fatal(err)
	////	}
	////}
	//pathes, err := repository.GetAllPathes()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, val := range pathes {
	//	fmt.Printf("Actual: %s; First: %s; Last: %s;\n", val.ActualPath, val.FirstCallPath, val.LastCallPath)
	//}
	//findedPath, isExist := repository.FindPathItemByFirstPath("FirstPath 3")
	//if !isExist {
	//	log.Println("FirstPath 3 not found")
	//} else {
	//	fmt.Printf("%#+v", findedPath)
	//}
	fmt.Print(generator.GenerateTableCode("brands", "Id", "Name", "Description"))
	cmd.Execute()
}

//
//type TestBrandTable struct {
//	Name    string
//	Columns TestBrandTableColumns
//}
//type TestBrandTableColumns struct {
//	Id          entities.Column
//	Name        entities.Column
//	Description entities.Column
//}
//
//func (t TestBrandTable) GetName() string {
//	return t.Name
//}
//func (t TestBrandTable) GetColumns() []entities.Column {
//	return []entities.Column{t.Columns.Id, t.Columns.Name, t.Columns.Description}
//}
//func (t TestBrandTable) GetColumnsCount() int {
//	return 3
//}
//func (t TestBrandTable) String() string {
//	return t.Name
//}
