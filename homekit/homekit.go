package homekit

import (
	"fmt"
	"net/url"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/chris-rock/homekit-fritz/setupcode"
	"github.com/sirupsen/logrus"
)

// HKConfig contains the configuration for the HomeKit service
type HKConfig struct {
	Pin     string `yaml:"pin"`
	SetupID string `yaml:"setupid"`
}

// FritzBoxConfig contains the confiratioin to access the Fritz!Box API
type FritzBoxConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Insecure bool   `yaml:"insecure"`
}

// GetFritzBoxURL parses the url and return its object
func (fbc *FritzBoxConfig) GetFritzBoxURL() *url.URL {
	url, err := url.Parse(fbc.URL)
	if err != nil {
		return nil
	}
	return url
}

// Config combines the Fritz!Box and the HomeKit configuration
type Config struct {
	HomeKit  *HKConfig       `yaml:"homekit"`
	FritzBox *FritzBoxConfig `yaml:"fritzbox"`
}

// Qrcode prints out the setup code based on the configuration
func Qrcode(hk *HKConfig) {
	xhmuri := setupcode.GenXhmURI(uint(accessory.TypeBridge), 0, hk.Pin, hk.SetupID)
	qrcode := setupcode.GenCliQRCode(xhmuri)
	fmt.Println(qrcode)
}

// Start is the HomeKit service that runs when you start `hkfritz serve`
func Start(config *Config) {

	// create fritzbox gateway
	fbBridge, err := CreateBridge(config.FritzBox)

	// read smart home devices
	hkDevices, err := ListHKDevices(config.FritzBox)

	// configure homekit service
	hcconfig := hc.Config{Pin: config.HomeKit.Pin, SetupId: config.HomeKit.SetupID}

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
