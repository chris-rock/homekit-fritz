package setupcode

import (
	"strings"

	"strconv"

	"github.com/sirupsen/logrus"
	qrcode "github.com/skip2/go-qrcode"
)

// GenXhmUri generates the xhm uri for that includes the pin and the setupid
// References are:
// - https://github.com/nfarina/homebridge/issues/1588#issuecomment-341158722
// - https://github.com/KhaosT/HAP-NodeJS/blob/67032e75b9f5f74993ad932c849d5bbb1937a097/lib/Accessory.js#L363
func GenXhmUri(category uint, hapType uint, pin string, setupId string) string {

	prefix := "X-HM://00"
	var payload uint64

	payload = 0
	cat := category << 31
	payload |= uint64(cat)

	ip := 1 << 28
	payload |= uint64(ip)

	u, err := strconv.ParseUint(pin, 10, 64)
	if err != nil {
		logrus.Error(err)
	}

	payload |= uint64(u)

	// covert to base 36
	s36 := strings.ToUpper(strconv.FormatUint(payload, 36))

	content := prefix + s36 + setupId
	return content
}

func GenCliQRCode(xhm string) string {
	var q *qrcode.QRCode
	q, err := qrcode.New(xhm, qrcode.Highest)
	if err != nil {
		logrus.Fatal(err)
		return ""
	}

	art := q.ToString(false)
	return art
}
