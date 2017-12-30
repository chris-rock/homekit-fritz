package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bpicode/fritzctl/logger"
	"github.com/pkg/errors"
)

// Config stores client configuration of your FRITZ!Box
type Config struct {
	*Net
	*Login
	*Pki
}

// Net wraps the protocol://host:port data to contact the FRITZ!Box.
type Net struct {
	Protocol string `json:"protocol"` // The protocol to use when communicating with the FRITZ!Box. "http" or "https".
	Host     string `json:"host"`     // Host name or ip address of the FRITZ!Box. In most home setups "fritz.box" can be used. Other possible formats: "192.168.2.200".
	Port     string `json:"port"`     // Port to use for the HTTP interface. Leave empty for default values.
}

// Login wraps the login data to be used by the client.
type Login struct {
	LoginURL string `json:"loginURL"` // The URL for the login negotiation.
	Username string `json:"username"` // Username to log in. In user-agnostic setups this can be left empty.
	Password string `json:"password"` // The password corresponding to the Username.
}

// Pki wraps the client-side certificate handling.
type Pki struct {
	SkipTLSVerify   bool   `json:"skipTlsVerify"`   // Skip TLS verification when using https.
	CertificateFile string `json:"certificateFile"` // Points to a certificate file (in PEM format) to verify the integrity of the FRITZ!Box.
}

// New creates a new Config by reading from a file given by the path.
func New(path string) (*Config, error) {
	logger.Debug("Reading config file", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot open configuration file '%s'", path)
	}
	conf := Config{}
	net := Net{}
	pki := Pki{}
	login := Login{}
	err = json.NewDecoder(file).Decode(&struct {
		*Net
		*Login
		*Pki
	}{&net, &login, &pki})
	conf.Pki = &pki
	conf.Login = &login
	conf.Net = &net
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse configuration file '%s'", path)
	}
	return &conf, nil
}

// GetLoginURL returns the URL that is queried for the login challenge
func (config *Config) GetLoginURL() string {
	return fmt.Sprintf("%s://%s:%s%s", config.Net.Protocol, config.Net.Host, config.Net.Port, config.Login.LoginURL)
}

// GetLoginResponseURL returns the URL that is queried for the login challenge
func (config *Config) GetLoginResponseURL(response string) string {
	return fmt.Sprintf("%s?response=%s&username=%s", config.GetLoginURL(), response, config.Login.Username)
}
