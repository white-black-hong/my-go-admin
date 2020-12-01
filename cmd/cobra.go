package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"my-go-admin/cmd/api"
	"my-go-admin/common/global"
	"os"
)

var rootCmd = &cobra.Command{
	Use:			"my-go-admin",
	Short:			"my-go-admin",
	SilenceUsage:	true,
	Long:			`my-go-admin`,
	Example:		`this is test`,
	Args: func(cmd *cobra.Command, args []string) error {
		//if len(args) < 1 {
		//	tip()
		//	return errors.New("requires at least one arg")
		//}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	// 必须要有否则 cobra 无法启动
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is a rootcmd")
	},
}

func tip()  {
	usageStr := `欢迎使用 my-go-admin ` + global.Version + ` 可以使用 ` + `-h` + ` 查看命令`
	usageStr1 := `也可以参考 http://doc.zhangwj.com/go-admin-site/guide/ksks.html 里边的【启动】章节`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init()  {
	rootCmd.AddCommand(api.StratCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("出错了")
		os.Exit(-1)
	}
}
