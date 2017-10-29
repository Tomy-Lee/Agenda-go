package cmd

import (
		"fmt"
		"io/ioutil"
		"log"
		"os"
	
	homedir 
		"github.com/spf13/cobra"
		"github.com/spf13/viper"
		"github.com/mitchellh/go-homedir"

)



/*
	Name: RootCmd
	Function: represents the base command 
 */
var RootCmd = &cobra.Command{
	Use:   "Agenda",
	Short: "A CLI meeting manager",
	Long: `Agenda supports different operation on meetings including register, create meeting, query and so on.
			It's a cooperation homework assignment for service computing.`,
}

/*
	Name: Execute
	Function: execute the root command
 */
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var File string
/*
	Name: initConfig
	Function: init the config
 */
func initConfig() {
	var db, _ = RootCmd.Flags().GetBool("debug")
	if !db {
		log.SetOutput(ioutil.Discard)
	}

	if File != "" {
		
		viper.SetConfigFile(File)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().BoolP("debug", "d", false, "display log message")
}
