package cmd

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var file string
var profile string
var region string
var output string
var cfgFile string
var service []string

var rootCmd = &cobra.Command{
	Use:   "iaws",
	Short: "AWS Insights CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("root called")
	},
}

func Execute() {
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "Output file")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "", "Use a specific profile from your credential file")
	rootCmd.PersistentFlags().StringVar(&region, "region", "ap-northeast-1", "The region to use. Overrides config/env settings")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "excel", "The formatting style for command output (excel, pdf, json, ...)")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "$HOME/.iaws.yaml", "Configuration file")

	rootCmd.Flags().StringArrayVar(&service, "service", nil, "Services(ec2,rds,s3,...). (default \"all\")")
	//rootCmd.MarkPersistentFlagRequired("region")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			rootCmd.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".iaws")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		rootCmd.Println("Using config file:", viper.ConfigFileUsed())
	}
}
