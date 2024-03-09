package key

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"keito/cmd"
	"keito/lib/keys"
	"keito/lib/util"
)

var keygen = &cobra.Command{
	Use:     "key",
	Aliases: []string{"k", "kg"},
	Short:   "Generate security key",
	Long:    `Generate security key.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := keys.Generate(viper.GetString(viperKeyAlgoParam), viper.GetInt(viperKeyLengthParam))
		if err != nil {
			fmt.Printf("\nerror: %s\n\n", err)
		} else {
			fmt.Printf("\n%s\n\n", key)
			if viper.GetBool(viperKeySaveParam) {
				util.SaveKeyConfig(key)
			}
		}
	},
}

const (
	cobraKeyLengthFlag  = "length"
	viperKeyLengthParam = "keygen.length"
	cobraKeyAlgoFlag    = "algo"
	viperKeyAlgoParam   = "keygen.algo"
	cobraKeySaveFlag    = "save"
	viperKeySaveParam   = "keygen.save"
)

func init() {
	keygen.Flags().IntP(cobraKeyLengthFlag, "l", 0, "key length")
	_ = viper.BindPFlag(viperKeyLengthParam, keygen.Flags().Lookup(cobraKeyLengthFlag))
	keygen.Flags().StringP(cobraKeyAlgoFlag, "a", "", "signature algorithm; supported values: hs256, hs384, hs512, rs256, rs384, rs512, es256, es384, es512, ps256, ps384, ps512")
	_ = viper.BindPFlag(viperKeyAlgoParam, keygen.Flags().Lookup(cobraKeyAlgoFlag))
	keygen.Flags().BoolP(cobraKeySaveFlag, "s", false, "save generated key to user home-located keito config")
	_ = viper.BindPFlag(viperKeySaveParam, keygen.Flags().Lookup(cobraKeySaveFlag))
	cmd.RootCmd.AddCommand(keygen)
}
