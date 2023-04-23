package diff

import (
	"errors"
	"sort"

	"github.com/zzzeep/diff-httpx/parser"
)

var ErrNotFound = errors.New("Resource was not found")

func GetChanges(oldRecords []parser.HttpxRecord, newRecords []parser.HttpxRecord) []Change {
	var changes []Change

	for _, nr := range newRecords {
		idx, err := FindIdxByInput(oldRecords, nr.Input)
		if errors.Is(err, ErrNotFound) {
			continue
		}
		or := oldRecords[idx]

		sort.Strings(or.A)
		sort.Strings(nr.A)

		if NeedsChange(or.A, nr.A) {
			newIps := getDifferece(or.A, nr.A)
			oldIps := getDifferece(nr.A, or.A)
			ch := Change{
				OldValue:   oldIps,
				NewValue:   newIps,
				ChangeType: ARecord,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Port, nr.Port) {
			ch := Change{
				OldValue:   or.Port,
				NewValue:   nr.Port,
				ChangeType: Port,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.StatusCode, nr.StatusCode) {
			ch := Change{
				OldValue:   or.StatusCode,
				NewValue:   nr.StatusCode,
				ChangeType: StatusCode,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Webserver, nr.Webserver) {
			ch := Change{
				OldValue:   or.Webserver,
				NewValue:   nr.Webserver,
				ChangeType: Webserver,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.ContentType, nr.ContentType) {
			ch := Change{
				OldValue:   or.ContentType,
				NewValue:   nr.ContentType,
				ChangeType: ContentType,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.ContentLength, nr.ContentLength) {
			ch := Change{
				OldValue:   or.ContentLength,
				NewValue:   nr.ContentLength,
				ChangeType: ContentLength,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Title, nr.Title) {
			ch := Change{
				OldValue:   or.Title,
				NewValue:   nr.Title,
				ChangeType: Title,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Hash.Body_md5, nr.Hash.Body_md5) {
			ch := Change{
				OldValue:   or.Hash.Body_md5,
				NewValue:   nr.Hash.Body_md5,
				ChangeType: BodyMD5,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Hash.Header_md5, nr.Hash.Header_md5) {
			ch := Change{
				OldValue:   or.Hash.Header_md5,
				NewValue:   nr.Hash.Header_md5,
				ChangeType: HeaderMD5,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
	}
	return changes
}

func FindIdxByInput(recs []parser.HttpxRecord, input string) (int, error) {
	for i, v := range recs {
		if v.Input == input {
			return i, nil
		}
	}
	return -1, ErrNotFound
}

func NeedsChange(oldV any, newV any) bool {

	if oldSli, ok := oldV.([]string); ok {
		newSli, _ := newV.([]string)
		return !testEqualSlices(oldSli, newSli)
	}

	if oldV != nil && oldV != newV {
		return true
	}
	return false
}

func testEqualSlices(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getDifferece(a []string, b []string) []string {
	ret := []string{}
	for _, v := range b {
		i := findStr(a, v)
		if i < 0 {
			ret = append(ret, v)
		}
	}
	return ret
}

func findStr(a []string, s string) int {
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1
}
