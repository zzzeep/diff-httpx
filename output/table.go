package output

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/zzzeep/diff-httpx/diff"
)

var NoColor bool

func PrintTable(changes []diff.Change) {
	if NoColor {
		text.DisableColors()
	}
	tbl := table.NewWriter()

	tbl.Style().Options.SeparateColumns = false
	tbl.Style().Options.SeparateRows = false
	tbl.Style().Options.DrawBorder = false
	tbl.SetOutputMirror(os.Stdout)

	for _, ch := range changes {
		if ch.ChangeType == diff.HeaderMD5 ||
			ch.ChangeType == diff.BodyMD5 ||
			ch.ChangeType == diff.ContentLength {
			continue
		}
		r := ch.ToTableRow()
		tbl.AppendRow(r)
	}

	tbl.Render()
}
