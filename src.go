//
//  go-webview-src-builder
//  Create a self-contained binary of HTML sources in a WebView
//
//  Copyright 2026, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package main

import (
	_ "embed"

	"github.com/nuxy/go-webview-src-builder/lib"
)

//go:embed src/index.html
var htmlDoc []byte

// Makefile linker variables
var Version string
var DevTools string

// Browser settings (defaults).
var settings = lib.BrowserSettings{
	Title:  "WebView Source",
	Height: 590,
	Width:  620,
	Resize: true,
	Debug: lib.StrToBool(DevTools),
}

// Let's get this party started.
func main() {
	browser := lib.NewBrowser(htmlDoc, settings)
	browser.Open()
}
