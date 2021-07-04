package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long: `Start api server

Start the API server. User the -h to check for available configuration flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called", viper.GetString("server.port"), viper.GetString("verbosity"))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringP("port", "p", "", "server port")
	serverCmd.Flags().VisitAll(configureViper("server"))
}
