package homekit

import (
	"fmt"
	"strconv"

	"github.com/bpicode/fritzctl/fritz"
	"github.com/brutella/hc/accessory"
	"github.com/sirupsen/logrus"
)

func NewHKSwitch(info accessory.Info, h fritz.HomeAuto, fdevice fritz.Device) (*accessory.Accessory, error) {
	logrus.Debugf("create new HomeKit outlet device %v", info)

	acc := accessory.NewOutlet(info)
	// add firmware informaation
	acc.Info.FirmwareRevision.SetValue(fdevice.Fwversion)

	// set initial state
	if s, err := strconv.ParseBool(fdevice.Switch.State); err == nil {
		acc.Outlet.On.SetValue(s)
		logrus.Debugf("initial state for %s %d", info.Name, s)
	}

	// add listener for device
	acc.Outlet.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			logrus.Printf("%s changed outlet to on", fdevice.Name)
			h.On(fdevice.Name)
		} else {
			logrus.Printf("%s changed outlet to off", fdevice.Name)
			h.Off(fdevice.Name)
		}
	})

	return acc.Accessory, nil
}

func NewHKThermostat(info accessory.Info, h fritz.HomeAuto, fdevice fritz.Device) (*accessory.Accessory, error) {
	logrus.Debugf("create new HomeKit thermostat device %v", info)

	// default values
	current := 20
	min := 8
	max := 28
	steps := 0.1

	acc := accessory.NewThermostat(info, float64(current), float64(min), float64(max), steps)

	// add firmware informaation
	acc.Info.FirmwareRevision.SetValue(fdevice.Fwversion)

	// read measured temprature
	if s, err := strconv.ParseFloat(fdevice.Thermostat.Measured, 32); err == nil {
		acc.Thermostat.CurrentTemperature.SetValue(s)
		logrus.Debugf("initial state for %s %d", info.Name, s)
	}

	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) {
		logrus.Printf("%s changed thermostat to on", fdevice.Name)
		h.Temp(temp, fdevice.Name)
	})

	return acc.Accessory, nil
}

func NewHomeKitDevice(h fritz.HomeAuto, fdevice fritz.Device) (*accessory.Accessory, error) {
	logrus.Infof("register %s", fdevice.Name)

	// create metadata for device
	info := accessory.Info{
		Name:         fdevice.Name,
		SerialNumber: fdevice.Identifier,
		Manufacturer: "AVM Computersysteme Vertriebs GmbH",
		Model:        fdevice.Productname,
	}

	if fdevice.IsSwitch() {
		return NewHKSwitch(info, h, fdevice)
	} else if fdevice.IsThermostat() {
		return NewHKThermostat(info, h, fdevice)
	} else {
		return nil, fmt.Errorf("do not support device %s", fdevice.Name)
	}
}
