package style

import (
	"net"

	"github.com/jedib0t/go-pretty/v6/text"
)

func StyleContentLength(v uint) string {
	return text.FgMagenta.Sprint(v)
}

func StyleStatusCode(v uint) string {
	if v >= 200 && v < 300 {
		return text.FgHiGreen.Sprint(v)
	} else if v >= 300 && v < 400 {
		return text.FgHiYellow.Sprint(v)
	} else {
		return text.FgHiRed.Sprint(v)
	}
}

func StyleWebServer(v string) string {
	return text.FgCyan.Sprint(v)
}

func StyleContentType(v string) string {
	return text.FgMagenta.Sprint(v)
}

func StyleARecord(ip net.IP) string {
	truncated := trunc(ip.String(), 20)
	if ip.IsPrivate() {
		return text.FgRed.Sprint(truncated)
	} else if ip.To4() != nil {
		return text.FgHiBlue.Sprint(truncated)
	} else {
		return text.FgBlue.Sprint(truncated)
	}
}

func StyleUrl(v string) string {
	return text.FgHiWhite.Sprint(v)
}

func StylePort(v string) string {
	return text.BgMagenta.Sprintf(":%s ", v)
}

func trunc(v string, i int) string {
	if len(v) > i {
		return v[:i] + ".."
	}
	return v
}
