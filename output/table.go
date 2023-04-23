package output

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/zzzeep/diff-httpx/change"
)

var Options DisplayOptions

func PrintTable(changes []change.Change) {
	Options.TestDefault()
	if Options.NoColor {
		text.DisableColors()
	}
	tbl := table.NewWriter()

	tbl.Style().Options.SeparateColumns = false
	tbl.Style().Options.SeparateRows = false
	tbl.Style().Options.DrawBorder = false
	tbl.SetOutputMirror(os.Stdout)

	for _, ch := range changes {

		if !Options.IPs && ch.ChangeType == change.IP {
			continue
		}
		if !Options.Port && ch.ChangeType == change.Port {
			continue
		}
		if !Options.Webserver && ch.ChangeType == change.Webserver {
			continue
		}
		if !Options.StatusCode && ch.ChangeType == change.StatusCode {
			continue
		}
		if !Options.Title && ch.ChangeType == change.Title {
			continue
		}
		if !Options.ContentType && ch.ChangeType == change.ContentType {
			continue
		}
		if !Options.ContentLength && ch.ChangeType == change.ContentLength {
			continue
		}
		if !Options.Hash && (ch.ChangeType == change.BodyMD5 || ch.ChangeType == change.HeaderMD5) {
			continue
		}
		r := ch.ToTableRow()
		tbl.AppendRow(r)
	}
	tbl.Render()
}
