package show

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/boyane126/bcpt/internal/bcptctl/cmd/util"
)

var Name string

func NewCmdShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "测试命令",
		Long:  "一个测试命令，用于测试是否使用正常",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(Run(args))
		},
	}

	cmd.Flags().StringVarP(&Name, "name", "n", "", "host 不能为空")

	return cmd
}

func Run(args []string) error {
	fmt.Println(args)
	return nil
}
