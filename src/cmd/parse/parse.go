package token

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"keito/cmd"
	"keito/lib/printer"
	"keito/lib/tokens"
)

const (
	cobraParseTokenFlag  = "token"
	viperParseTokenParam = "parse.token"

	cobraParseKeyFlag  = "key"
	viperParseKeyParam = "parse.key"
)

var tokengen = &cobra.Command{
	Use:     "parse",
	Aliases: []string{"p", "ps"},
	Short:   "Parse JWT token",
	Long:    `Parse JWT token.`,
	Run: func(cmd *cobra.Command, args []string) {
		parsed, verified, err := tokens.Parse(
			viper.GetString(viperParseTokenParam),
			viper.GetString(viperParseKeyParam),
		)
		if err != nil {
			fmt.Printf("\nerror: %s\n\n", err)
		} else {
			fmt.Println()
			fmt.Println("Claims:")
			printer.PrintMapSorted(parsed, "  ")
			if verified {
				fmt.Println("\nSignature verified!")
			}
			fmt.Println()
		}
	},
}

func init() {

	tokengen.Flags().StringP(cobraParseTokenFlag, "t", "", "JWT token to be parsed")
	_ = viper.BindPFlag(viperParseTokenParam, tokengen.Flags().Lookup(cobraParseTokenFlag))

	tokengen.Flags().StringP(cobraParseKeyFlag, "k", "", "JWT token signing key for token signature verification (optional)")
	_ = viper.BindPFlag(viperParseKeyParam, tokengen.Flags().Lookup(cobraParseKeyFlag))

	cmd.RootCmd.AddCommand(tokengen)
}
