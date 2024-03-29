package cmd

import (
	"fmt"
	"os"

	"github.com/nishi-yuki/knetlogin/knet"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	id      string
	pass    string
)

var rootCmd = &cobra.Command{
	Use:     "knetlogin",
	Short:   "K**net auto login tool",
	Version: "alpha 0.10",
	RunE: func(cmd *cobra.Command, args []string) error {
		id = viper.GetString("id")
		pass = viper.GetString("pass")

		fmt.Println("start login process.")
		err := knet.Login(id, pass)
		if err == nil {
			fmt.Println("login succeeded.")
		} else {
			fmt.Println("login failed.")
		}
		return err
	},
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/knetlogin.yaml)")
	rootCmd.PersistentFlags().StringP("id", "i", "", "your user id")
	rootCmd.PersistentFlags().StringP("pass", "p", "", "your password")
	viper.BindPFlag("id", rootCmd.PersistentFlags().Lookup("id"))
	viper.BindPFlag("pass", rootCmd.PersistentFlags().Lookup("pass"))
}
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("knetlogin")
		viper.AddConfigPath("$HOME/.config")
		viper.AddConfigPath(".")

	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
