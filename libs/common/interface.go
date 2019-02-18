package sbi

// Config define the raw data of plugin
type Config struct {
	Dev bool

	PKG     string
	Comment string

	Algorithm string
}

// Plugin is the interface of plugin
type Plugin interface {
	Encode()
	Example()
	Decode()
}
