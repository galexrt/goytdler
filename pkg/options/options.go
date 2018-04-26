package options

type CmdLineOpts struct {
	LogLevel      string
	ListenAddress string
	YoutubeDLPath string
	OutputPath    string
}

var Opts CmdLineOpts
