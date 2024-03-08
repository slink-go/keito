package token

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"keito/cmd"
	"keito/lib/tokens"
)

const (
	cobraTokenFlag  = "token"
	viperTokenParam = "parse.token"
)

var tokengen = &cobra.Command{
	Use:     "parse",
	Aliases: []string{"p", "ps"},
	Short:   "Parse JWT token",
	Long:    `Parse JWT token.`,
	Run: func(cmd *cobra.Command, args []string) {
		parsed, err := tokens.Parse(
			viper.GetString(viperTokenParam),
		)
		if err != nil {
			fmt.Printf("\nerror: %s\n\n", err)
		} else {
			fmt.Println()
			for k, v := range parsed {
				fmt.Printf("%s = %v\n", k, v)
			}
			fmt.Println()
		}
	},
}

func init() {

	tokengen.Flags().StringP(cobraTokenFlag, "t", "", "JWT token to be parsed")
	_ = viper.BindPFlag(viperTokenParam, tokengen.Flags().Lookup(cobraTokenFlag))

	cmd.RootCmd.AddCommand(tokengen)
}
