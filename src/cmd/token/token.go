package token

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"keito/cmd"
	"keito/lib/tokens"
)

const (
	cobraTokenAlgoFlag  = "algo"
	viperTokenAlgoParam = "token.algo"

	cobraTokenKeyFlag  = "key"
	viperTokenKeyParam = "token.key"

	cobraTokenIssuerFlag  = "issuer"
	viperTokenIssuerParam = "token.issuer"

	cobraTokenSubjectFlag  = "subject"
	viperTokenSubjectParam = "token.subject"

	cobraTokenDurationFlag  = "duration"
	viperTokenDurationParam = "token.duration"

	cobraTokenClaimsFlag  = "claims"
	viperTokenClaimsParam = "token.claims"

	cobraTokenOneTimeFlag  = "onetime"
	viperTokenOneTimeParam = "token.onetime"
)

var tokengen = &cobra.Command{
	Use:     "token",
	Aliases: []string{"t", "tk"},
	Short:   "Generate JWT token",
	Long:    `Generate JWT token.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := tokens.Generate(
			viper.GetString(viperTokenAlgoParam),
			viper.GetString(viperTokenKeyParam),
			viper.GetString(viperTokenIssuerParam),
			viper.GetString(viperTokenSubjectParam),
			viper.GetString(viperTokenDurationParam),
			viper.GetString(viperTokenClaimsParam),
			viper.GetBool(viperTokenOneTimeParam),
		)
		if err != nil {
			fmt.Printf("\nerror: %s\n\n", err)
		} else {
			fmt.Printf("\n%s\n\n", key)
		}
	},
}

func init() {

	tokengen.Flags().StringP(cobraTokenAlgoFlag, "a", "", "key signing algorithm (hs256, hs384, hs512)")
	_ = viper.BindPFlag(viperTokenAlgoParam, tokengen.Flags().Lookup(cobraTokenAlgoFlag))

	tokengen.Flags().StringP(cobraTokenKeyFlag, "k", "", "signature key (optional, if not set key will be automatically generated)")
	_ = viper.BindPFlag(viperTokenKeyParam, tokengen.Flags().Lookup(cobraTokenKeyFlag))

	tokengen.Flags().StringP(cobraTokenIssuerFlag, "i", "", "token issuer")
	_ = viper.BindPFlag(viperTokenIssuerParam, tokengen.Flags().Lookup(cobraTokenIssuerFlag))

	tokengen.Flags().StringP(cobraTokenSubjectFlag, "s", "", "token subject")
	_ = viper.BindPFlag(viperTokenSubjectParam, tokengen.Flags().Lookup(cobraTokenSubjectFlag))

	tokengen.Flags().StringP(cobraTokenDurationFlag, "d", "", "token duration (i.e. 15m, 1h, 3d)")
	_ = viper.BindPFlag(viperTokenDurationParam, tokengen.Flags().Lookup(cobraTokenDurationFlag))

	tokengen.Flags().StringP(cobraTokenClaimsFlag, "c", "", "token claims (comma-separated key=value pairs)")
	_ = viper.BindPFlag(viperTokenClaimsParam, tokengen.Flags().Lookup(cobraTokenClaimsFlag))

	tokengen.Flags().BoolP(cobraTokenOneTimeFlag, "o", false, "generate 'jti' claim for one-time-use token")
	_ = viper.BindPFlag(viperTokenOneTimeParam, tokengen.Flags().Lookup(cobraTokenOneTimeFlag))

	cmd.RootCmd.AddCommand(tokengen)
}
