# Homekit Fritz - HomeKit Bridge for Fritz!Box

This project adds the missing feature of a HomeKit Bridge to all Fritz!Box smart home devices.

Note: This is a hobby project and was stated by a few questions:
- How does homekit work?
- What is required to run a secure environment for smart home devices?
- Can everything on raspberry pi 1
- How can I manage my devices with Siri

## Features

- no need for exposing Fritz!Box to internet
- easy setup with HomeKit setup codes

## Tested Devices

- FRITZ!DECT 300 Thermostats
- FRITZ!DECT 200 Switches

## Getting Started

```
# asks for the fritzbox credentials
hkfritz configure

# starts the homekit bridge service
hkfritz serve

# cli to reprint the qr code
hkfritz setupcode
```

## Limitations

- no bi-directonal sync between Fritzbox and Homekit
- changes of names in Home app confuses the Gateway

## Kudos

This project is built ontop of the great libraries:

- https://github.com/bpicode/fritzctl
- https://github.com/brutella/hc

## Related projects

[Homebridge](https://github.com/nfarina/homebridge) with its [Fritz!Box Plugin](https://github.com/andig/homebridge-fritz) provide similar functionality.

## Author

- Christopoh Hartmann

