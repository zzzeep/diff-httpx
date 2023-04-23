package diff

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/zzzeep/diff-httpx/style"
)

type ChangeType string

const (
	ARecord       ChangeType = "A Record"
	Port          ChangeType = "Port"
	Webserver     ChangeType = "Web Server"
	StatusCode    ChangeType = "Status Code"
	ContentLength ChangeType = "Content Length"
	ContentType   ChangeType = "Content Type"
	Title         ChangeType = "Title"
	BodyMD5       ChangeType = "Body MD5"
	HeaderMD5     ChangeType = "Header MD5"
)

type Change struct {
	OldValue   any
	NewValue   any
	ChangeType ChangeType
	Url        string
}

func (ch *Change) Print() {
	fmt.Println(ch.OldValue, ch.NewValue, ch.ChangeType, ch.Url)
}

func (ch *Change) ToTableRow() table.Row {
	r := table.Row{}

	oldValFormatted := FormatValue(ch.OldValue, ch.ChangeType)
	newValFormatted := FormatValue(ch.NewValue, ch.ChangeType)

	r = append(r, oldValFormatted)
	r = append(r, ">")
	r = append(r, newValFormatted)
	r = append(r, ch.ChangeType)
	r = append(r, style.StyleUrl(ch.Url))
	return r
}

func FormatValue(v any, vType ChangeType) string {
	switch vType {
	case ContentLength:
		return style.StyleContentLength(v.(uint))
	case ContentType:
		return style.StyleContentType(v.(string))
	case StatusCode:
		return style.StyleStatusCode(v.(uint))
	case Webserver:
		return style.StyleWebServer(v.(string))
	case Port:
		return style.StylePort(v.(string))
	case ARecord:
		return style.StyleARecord(v.([]string))
	default:
		vStr := v.(string)
		return vStr
	}
}
