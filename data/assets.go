/*
Copyright 2020 Alexander Trost <galexrt@googlemail.com>

This work is licensed under the terms of the MIT license.
For a copy, see <https://opensource.org/licenses/MIT>.
*/

package data

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetse48cd5158def07113582a8696d146a718b0c2911 = "<!DOCTYPE html>\n<html>\n    <head>\n        <title>goytdler - Simple webinterface to use youtube-dl</title>\n    </head>\n    <body>\n        <h1>goytdler</h1>\n        <p>Simple webinterface to use youtube-dl.</p>\n        <hr>\n        <form method=\"POST\" action=\"{{ .RoutesBasePath }}download\">\n            <input type=\"text\" name=\"url\" placeholder=\"Youtube-DL compatible Link\" />\n            <input type=\"submit\" value=\"Submit\" />\n        </form>\n    </body>\n</html>\n"
var _Assets630f88879f98b6e0f05588023557b3717e54a9ae = "<!DOCTYPE html>\n<html>\n    <head>\n        <title>goytdler - Simple webinterface to use youtube-dl</title>\n    </head>\n    <body>\n        <h1>goytdler</h1>\n        <p>Simple webinterface to use youtube-dl.</p>\n        <hr>\n        <h2>Error</h2>\n        <p>{{ . }}</p>\n    </body>\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"index.tmpl", "error.tmpl"}}, map[string]*assets.File{
	"/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1524760173, 1524760173068371739),
		Data:     nil,
	}, "/templates/index.tmpl": &assets.File{
		Path:     "/templates/index.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1524825723, 1524825723310010286),
		Data:     []byte(_Assetse48cd5158def07113582a8696d146a718b0c2911),
	}, "/templates/error.tmpl": &assets.File{
		Path:     "/templates/error.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1524762038, 1524762038861083087),
		Data:     []byte(_Assets630f88879f98b6e0f05588023557b3717e54a9ae),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1524825340, 1524825340852696592),
		Data:     nil,
	}}, "")
