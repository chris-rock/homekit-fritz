package homekit

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bpicode/fritzctl/config"
	"github.com/bpicode/fritzctl/fritz"
	"github.com/brutella/hc/accessory"
	"github.com/sirupsen/logrus"
)

func CreateBridge(fbConfig *FritzBoxConfig) (*accessory.Accessory, error) {

	fburl := fbConfig.GetFritzBoxURL()
	if fburl == nil {
		return nil, fmt.Errorf("no Fritz!Box url provided")
	}

	// read fritzbox info
	cfg := &config.Config{
		Net:   &config.Net{Protocol: fburl.Scheme, Host: fburl.Host, Port: ""},
		Login: &config.Login{LoginURL: "/login_sid.lua", Username: fbConfig.Username, Password: fbConfig.Password},
		Pki:   &config.Pki{SkipTLSVerify: true, CertificateFile: ""},
	}

	logrus.Debug(cfg)

	// TODO: support tls
	transport := &http.Transport{}
	httpClient := &http.Client{Transport: transport}
	c := &fritz.Client{Config: cfg, HTTPClient: httpClient}
	err := c.Login()
	if err != nil {
		return nil, err
	}
	f := fritz.NewInternal(c)
	info, err := f.BoxInfo()
	if err != nil {
		return nil, err
	}
	logrus.Infof("FritzBox: %s", info.Model.Name)

	// TODO: extract to sanitize name function
	fbName := strings.Replace(info.Model.Name, "!", "", -1)

	fritzBoxInfo := accessory.Info{
		Name:         fbName,
		Manufacturer: "AVM Computersysteme Vertriebs GmbH",
		Model:        "Fritz!Box",
	}
	fritzboxAccessory := accessory.New(fritzBoxInfo, accessory.TypeBridge)
	return fritzboxAccessory, nil

}
