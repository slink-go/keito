package cmd

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

//go:embed logo.txt
var logo string

var RootCmd = &cobra.Command{
	Use:   filepath.Base(os.Args[0]),
	Short: "Key/token tool",
	Long: logo + ` 

Key & token generator

    - generate cryptographically secure random keys
    - generate JWT tokens, signed with previously generated keys
    - parse claims from JWT tokens

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

func init() {
	RootCmd.SetUsageTemplate(CustomUsageTemplate)
}
