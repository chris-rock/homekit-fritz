package homekit

import (
	"fmt"
	"net/url"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/chris-rock/homekit-fritz/setupcode"
	"github.com/sirupsen/logrus"
)

// Configuration
type HomeKitConfig struct {
	Pin     string `yaml:"pin"`
	SetupId string `yaml:"setupid"`
}

type FritzBoxConfig struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Insecure bool   `yaml:"insecure"`
}

func (fbc *FritzBoxConfig) GetFritzBoxURL() *url.URL {
	url, err := url.Parse(fbc.Url)
	if err != nil {
		return nil
	}
	return url
}

type Config struct {
	HomeKit  *HomeKitConfig  `yaml:"homekit"`
	FritzBox *FritzBoxConfig `yaml:"fritzbox"`
}

func Qrcode(hk *HomeKitConfig) {
	xhmuri := setupcode.GenXhmUri(uint(accessory.TypeBridge), 0, hk.Pin, hk.SetupId)
	qrcode := setupcode.GenCliQRCode(xhmuri)
	fmt.Println(qrcode)
}

// Service Implementation
func Start(config *Config) {

	// create fritzbox gateway
	fbBridge, err := CreateBridge(config.FritzBox)

	// read smart home devices
	hkDevices, err := ListHKDevices(config.FritzBox)

	// configure homekit service
	hcconfig := hc.Config{Pin: config.HomeKit.Pin, SetupId: config.HomeKit.SetupId}

	// create fritzbox as bridge device with all attached home kit devices
	t, err := hc.NewIPTransport(hcconfig, fbBridge, hkDevices...)
	if err != nil {
		logrus.Panic(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
