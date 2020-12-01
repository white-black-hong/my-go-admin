package api

import (
	"github.com/spf13/cobra"
	"my-go-admin/tools/config"
)

var (
	configYml	string
	port		string
	mode		string
	traceStart	bool
	StratCmd	= &cobra.Command{
		Use:			"server",
		Short:			"Start API server",
		Example:		"my-go-admin server -c config/setting.dev.xml",
		SilenceUsage:	true,
	PreRun: func(cmd *cobra.Command, args []string) {
		setup()
	},
	}
)


func init()  {
	StratCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/setting.dev.xml","Start server with provided configuration file")

	// 注册路由

}

func setup()  {

	// 1. 读取配置
	config.Setup(configYml)
	// 2. 设置日志


}
