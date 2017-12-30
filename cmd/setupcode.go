package cmd

import (
	"fmt"

	"github.com/chris-rock/homekit-fritz/homekit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setupcodeCmd represents the setupcode command
var setupcodeCmd = &cobra.Command{
	Use:   "setupcode",
	Short: "Print HomeKit SetupCode",
	Long:  `This command prints the setup code that can be used within the Home app`,
	Run: func(cmd *cobra.Command, args []string) {
		hk := &homekit.HKConfig{
			Pin:     viper.GetString("homekit.pin"),
			SetupID: viper.GetString("homekit.setupid"),
		}

		fmt.Println("HomeKit setup qr code:")
		homekit.Qrcode(hk)
		fmt.Printf("HomeKit setup code: %s\n", hk.Pin)
	},
}

func init() {
	rootCmd.AddCommand(setupcodeCmd)
}
