package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-api",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())

	lvl, err := log.ParseLevel(viper.GetString("verbosity"))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed to parse log level")
	}
	log.SetLevel(lvl)
}

func init() {
	// Change your env prefix here
	viper.SetDefault("envPrefix", "GO_API")
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("verbosity", "v", "", "set log verbosity")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(viper.GetString("envPrefix"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.WithFields(log.Fields{
			"file": viper.ConfigFileUsed(),
		}).Info("Using config file")
	}
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		panic(err)
	}
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		panic(err)
	}
}

func configureViper(prefix string) (fn func(*pflag.Flag)) {
	return func(flag *pflag.Flag) {
		prefixedName := strings.Join([]string{prefix, flag.Name}, ".")
		if err := viper.BindPFlag(prefixedName, flag); err != nil {
			panic(err)
		}
		env := strings.Join([]string{viper.GetString("envPrefix"), strings.ReplaceAll(prefixedName, ".", "_")}, "_")
		err := viper.BindEnv(prefixedName, strings.ToUpper(env))
		if err != nil {
			log.WithFields(log.Fields{
				"env":   env,
				"error": err,
			}).Error("Failed to bind env var")
		}
	}
}
