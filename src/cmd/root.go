package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

//var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "kt",
	Short: "Key/token tool",
	Long: ` 
    key/token generator
    - generate cryptographically secure random keys
    - generate JWT tokens
    `,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		panic(fmt.Errorf("command error: %s", err.Error()))
	}
}

const CustomUsageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if and (ne .Name "completion") .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

const (
// ViperOutputNoHeaderFlag = "output.common.noheader"
// ViperOutputFormatFlag   = "output.common.format"
)

func init() {
	cobra.OnInitialize(initConfig)

	//RootCmd.PersistentFlags().BoolP("noheader", "n", false, "don't print header")
	//_ = viper.BindPFlag(ViperOutputNoHeaderFlag, RootCmd.PersistentFlags().Lookup("noheader"))
	//RootCmd.PersistentFlags().StringP("format", "o", "", "output format")
	//_ = viper.BindPFlag(ViperOutputFormatFlag, RootCmd.PersistentFlags().Lookup("format"))

	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rptcli)")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("KT")
	//viper.SetConfigType("yaml")
	RootCmd.SetUsageTemplate(CustomUsageTemplate)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	//if cfgFile != "" {
	//	// Use config file from the flag.
	//	viper.SetConfigFile(cfgFile)
	//} else {
	//	// Find home directory.
	//	home, err := os.UserHomeDir()
	//	cobra.CheckErr(err)
	//
	//	// Search config in home directory with name ".rptcli" (without extension).
	//	viper.AddConfigPath(home)
	//	viper.SetConfigName(".rptcli")
	//}

	// If a config file is found, read it in.
	//if err := viper.ReadInConfig(); err == nil {
	//	//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	//	//fmt.Fprintln(os.Stderr, viper.AllKeys())
	//} else {
	//	fmt.Fprintln(os.Stderr, "Error, reading config file: ", err)
	//}

	viper.AutomaticEnv() // read in environment variables that match

	//for _, v := range viper.AllKeys() {
	//	fmt.Println(v, viper.Get(v))
	//}
}
