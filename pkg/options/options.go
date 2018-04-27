package options

type CmdLineOpts struct {
	LogLevel       string
	ListenAddress  string
	YoutubeDLPath  string
	OutputPath     string
	RoutesBasePath string
}

var Opts CmdLineOpts
