package diff

import (
	"bytes"
	"errors"
	"net"
	"sort"

	"github.com/zzzeep/diff-httpx/change"
	"github.com/zzzeep/diff-httpx/output"
	"github.com/zzzeep/diff-httpx/parser"
)

var ErrNotFound = errors.New("Resource was not found")

func GetChanges(oldRecords []parser.HttpxRecord, newRecords []parser.HttpxRecord) []change.Change {
	var changes []change.Change

	for _, nr := range newRecords {
		if output.Options.FilterCode > 0 &&
			output.Options.FilterCode != nr.StatusCode {
			continue
		}

		idx, err := FindIdxByInput(oldRecords, nr.Input)
		if errors.Is(err, ErrNotFound) {
			continue
		}
		or := oldRecords[idx]

		oldIps := parseIPs(or.A)
		newIps := parseIPs(nr.A)

		if NeedsChange(or.A, nr.A) {
			newIPs := getDifferece(oldIps, newIps)
			oldIPs := getDifferece(newIps, oldIps)

			for i := range newIPs {
				ch := change.Change{
					OldValue:   getIPsafe(oldIPs, i),
					NewValue:   getIPsafe(newIPs, i),
					ChangeType: change.IP,
					Url:        nr.Url,
				}
				changes = append(changes, ch)
			}
		}
		if NeedsChange(or.Port, nr.Port) {
			ch := change.Change{
				OldValue:   or.Port,
				NewValue:   nr.Port,
				ChangeType: change.Port,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.StatusCode, nr.StatusCode) {
			ch := change.Change{
				OldValue:   or.StatusCode,
				NewValue:   nr.StatusCode,
				ChangeType: change.StatusCode,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Webserver, nr.Webserver) {
			ch := change.Change{
				OldValue:   or.Webserver,
				NewValue:   nr.Webserver,
				ChangeType: change.Webserver,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.ContentType, nr.ContentType) {
			ch := change.Change{
				OldValue:   or.ContentType,
				NewValue:   nr.ContentType,
				ChangeType: change.ContentType,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.ContentLength, nr.ContentLength) {
			ch := change.Change{
				OldValue:   or.ContentLength,
				NewValue:   nr.ContentLength,
				ChangeType: change.ContentLength,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Title, nr.Title) {
			ch := change.Change{
				OldValue:   or.Title,
				NewValue:   nr.Title,
				ChangeType: change.Title,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Hash.Body_md5, nr.Hash.Body_md5) {
			ch := change.Change{
				OldValue:   or.Hash.Body_md5,
				NewValue:   nr.Hash.Body_md5,
				ChangeType: change.BodyMD5,
				Url:        nr.Url,
			}
			changes = append(changes, ch)
		}
		if NeedsChange(or.Hash.Header_md5, nr.Hash.Header_md5) {
			ch := change.Change{
				OldValue:   or.Hash.Header_md5,
				NewValue:   nr.Hash.Header_md5,
				ChangeType: change.HeaderMD5,
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

func parseIPs(ips []string) []net.IP {
	var realIPs []net.IP
	for _, ip := range ips {
		realIPs = append(realIPs, net.ParseIP(ip))
	}
	sort.Slice(realIPs, func(i, j int) bool {
		return bytes.Compare(realIPs[i], realIPs[j]) < 0
	})
	return realIPs
}

func getDifferece(a []net.IP, b []net.IP) []net.IP {
	ret := []net.IP{}
	for _, v := range b {
		i := findIP(a, v)
		if i < 0 {
			ret = append(ret, v)
		}
	}
	return ret
}

func findIP(a []net.IP, b net.IP) int {
	for i, v := range a {
		if bytes.Compare(v, b) == 0 {
			return i
		}
	}
	return -1
}

func getIPsafe(ips []net.IP, i int) net.IP {
	if len(ips) <= i {
		return nil
	}
	return ips[i]
}
