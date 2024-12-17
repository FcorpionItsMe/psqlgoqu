package mktab

import (
	"github.com/FcorpionItsMe/psqlgoqu/internal/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var Path string
var Name string

var mktabCommand = &cobra.Command{
	Use: "mktab",

	Short: "Use mktab to generate tables(models) for queries.",
	Long: "Use mktab to generate tables(models) for queries. " +
		"path - path where table will be save" +
		"name - name of the table" +
		"args - list of columns",
	Run: func(cmd *cobra.Command, args []string) {
		projectPath, err := utils.GetCurrentDirPath()
		if err != nil {
			log.Fatal(err)
			return
		}
		if Path == "" {
			log.Fatal("Path is empty!")
			return
		}
		actualPath := projectPath + "\\" + Path
		stat, err := os.Stat(actualPath)
		if err != nil {
			log.Fatal(err)
			return
		}
		if !stat.IsDir() {
			log.Fatal("Path is not a directory!")
		}
		
	},
}

func InitCommand(root *cobra.Command) {
	root.AddCommand(mktabCommand)
	mktabCommand.Flags().StringVarP(&Path, "path", "p", "", "Path where the model will be created.")
	err := mktabCommand.MarkFlagRequired("path")
	if err != nil {
		log.Fatal(err)
		return
	}
	mktabCommand.Flags().StringVarP(&Name, "name", "n", "", "Name of the model that will be created.")
	err = mktabCommand.MarkFlagRequired("name")
	if err != nil {
		log.Fatal(err)
		return
	}
}
