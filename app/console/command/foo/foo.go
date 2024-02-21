package foo

import (
	"bbs/app/provider/qa"
	"bbs/app/provider/user"
	"fmt"
	"github.com/gohade/hade/framework/cobra"
	"github.com/gohade/hade/framework/contract"
)

var ModelInitCommand = &cobra.Command{
	Use:   "migrate",
	Short: "初始化model表",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()

		ormService := container.MustMake(contract.ORMKey).(contract.ORMService)
		db, err := ormService.GetDB()
		if err != nil {
			return err
		}
		if err := db.AutoMigrate(&user.User{}, &qa.Answer{}, &qa.Question{}); err != nil {
			return err
		}
		fmt.Println("migrate all table success")

		return nil
	},
}
