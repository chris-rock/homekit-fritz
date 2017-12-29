package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/chris-rock/homekit-fritz/homekit"
	yaml "gopkg.in/yaml.v2"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Command to configure the Fritz!Box credentials",
	Long:  `Prompts the Fritz!Box credentials and generates the configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		// initialize config
		cfg := &homekit.Config{
			FritzBox: &homekit.FritzBoxConfig{},
			HomeKit:  &homekit.HomeKitConfig{},
		}

		validateURL := func(input string) error {
			_, err := url.Parse(viper.GetString("fritzbox.url"))
			if err != nil {
				return errors.New("Invalid URL")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Frtz!Box Url",
			Validate: validateURL,
			Default:  "http://fritz.box",
		}
		fbURL, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		cfg.FritzBox.Url = fbURL

		prompt = promptui.Prompt{
			Label:     "Self-signed certificate",
			IsConfirm: true,
		}
		fbInsecure, _ := prompt.Run()
		if fbInsecure == "y" {
			cfg.FritzBox.Insecure = true
		}

		prompt = promptui.Prompt{
			Label: "Frtz!Box username",
		}
		fbUsername, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		cfg.FritzBox.Username = fbUsername

		prompt = promptui.Prompt{
			Label: "Frtz!Box password",
			Mask:  '*',
		}
		fbPassword, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		cfg.FritzBox.Password = fbPassword

		// generate PIN and Setup Code
		cfg.HomeKit.Pin = fmt.Sprintf("%08d", rand.Intn(100000000))
		// generate setup id
		cfg.HomeKit.SetupId = randStringBytesRmndr(4)

		d, err := yaml.Marshal(cfg)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// write configuration to file
		path := filepath.Join(userHomeDir(), ".hk-fritz.yml")
		fmt.Printf("write configuration to %s", path)

		err = ioutil.WriteFile(path, d, 0644)
		if err != nil {
			log.Fatalf("Could not write configuration file: %v", err)
		}
	},
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
