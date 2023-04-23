package style

import (
	"net"
	"strings"

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

func StyleARecord(v []string) string {
	for i := range v {
		ip := net.ParseIP(v[i])
		output := trunc(v[i], 20)

		if ip.IsPrivate() {
			v[i] = text.FgRed.Sprint(output)
		} else if ip.To4() != nil {
			v[i] = text.FgHiBlue.Sprint(output)
		} else {
			v[i] = text.FgBlue.Sprint(output)
		}
	}
	return strings.Join(v, "\n")
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
