package setpath

import (
	"github.com/FcorpionItsMe/psqlgoqu/internal/repository/json_repository"
	"github.com/FcorpionItsMe/psqlgoqu/internal/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var Path string
var repos *json_repository.Repository

var setpathCommand = &cobra.Command{
	Use:   "setpath",
	Short: "Use setpath to set path where tables(models) will be save.",
	Long:  "Use setpath to set path where tables(models) will be save.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("Args cannot be empty! ")
			return
		}
		if len(args) > 1 {
			log.Println("Too many arguments! ")
			return
		}
		arg := args[0]
		rootDir, err := utils.GetCurrentDirPath()
		if err != nil {
			log.Println(err)
		}
		actualPath := rootDir + "\\" + arg
		stat, err := os.Stat(actualPath)
		if err != nil {
			log.Println(err)
			return
		}
		if !stat.IsDir() {
			log.Println("Not a directory!")
		}
		//pathItem, isExist := repos.FindPathItemByActualPath(actualPath)
		//if isExist {
		//
		//}
	},
}

func InitCommand(root *cobra.Command, rep *json_repository.Repository) {
	repos = rep
	root.AddCommand(setpathCommand)
}
