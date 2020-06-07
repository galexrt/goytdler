/*
Copyright 2020 Alexander Trost <galexrt@googlemail.com>

This work is licensed under the terms of the MIT license.
For a copy, see <https://opensource.org/licenses/MIT>.
*/

package options

type CmdLineOpts struct {
	LogLevel       string
	ListenAddress  string
	YoutubeDLPath  string
	OutputPath     string
	RoutesBasePath string
}

var Opts CmdLineOpts
