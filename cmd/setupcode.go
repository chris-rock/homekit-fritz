package cmd

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/chris-rock/homekit-fritz/setupcode"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setupcodeCmd represents the setupcode command
var setupcodeCmd = &cobra.Command{
	Use:   "setupcode",
	Short: "Print HomeKit SetupCode",
	Long:  `This command prints the setup code that can be used within the Home app`,
	Run: func(cmd *cobra.Command, args []string) {
		pin := viper.GetString("homekit.pin")
		setupid := viper.GetString("homekit.setupid")
		xhmuri := setupcode.GenXhmUri(uint(accessory.TypeBridge), 0, pin, setupid)
		qrcode := setupcode.GenCliQRCode(xhmuri)
		fmt.Println("HomeKit setup qr code:")
		fmt.Println(qrcode)
		fmt.Printf("HomeKit setup code: %s\n", pin)
	},
}

func init() {
	rootCmd.AddCommand(setupcodeCmd)
}
