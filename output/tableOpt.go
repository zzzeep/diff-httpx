package output

type DisplayOptions struct {
	FilterCode uint
	NoColor    bool
	NoTrunc    bool

	IPs           bool
	Port          bool
	Webserver     bool
	StatusCode    bool
	Title         bool
	ContentType   bool
	ContentLength bool
	Hash          bool
}

func (d *DisplayOptions) TestDefault() {
	if !(d.IPs || d.Port || d.Webserver ||
		d.StatusCode || d.Title) {
		d.IPs = true
		d.Port = true
		d.Webserver = true
		d.StatusCode = true
		d.Title = true
	}
}
