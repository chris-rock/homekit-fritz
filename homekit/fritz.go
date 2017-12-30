package homekit

import (
	"github.com/bpicode/fritzctl/fritz"
	"github.com/brutella/hc/accessory"
	"github.com/sirupsen/logrus"
)

// ListHKDevices reads all smart home devices from Fritz!Box and maps them
// to HomeKit devices
func ListHKDevices(fbConfig *FritzBoxConfig) ([]*accessory.Accessory, error) {

	h := fritz.NewHomeAuto(
		fritz.URL(fbConfig.GetFritzBoxURL()),
		fritz.SkipTLSVerify(),
		fritz.Credentials(fbConfig.Username, fbConfig.Password),
	)

	err := h.Login()
	if err != nil {
		return nil, err
	}

	fdevices, err := h.List()
	if err != nil {
		return nil, err
	}

	// map switches to HomeKit
	hkDevices := make([]*accessory.Accessory, 0)
	for _, fdevice := range fdevices.Devices {
		s, err := NewHomeKitDevice(h, fdevice)

		if err == nil {
			hkDevices = append(hkDevices, s)
		} else {
			logrus.Errorf("do not understand the device %s", fdevice.Name)
			// ignore error, to continue with the devices we understand
		}
	}
	return hkDevices, nil
}
