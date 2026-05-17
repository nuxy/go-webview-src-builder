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
	"embed"

	"github.com/nuxy/go-webview-src-builder/lib"
)

//go:embed all:src
var content embed.FS

// Makefile linker variables
var Version string
var DevTools string

// Browser settings (defaults).
var settings = lib.BrowserSettings{
	Title:  "WebView Source",
	Height: 590,
	Width:  620,
	Resize: true,
	Debug:  lib.StrToBool(DevTools),
}

// Let's get this party started.
func main() {
	htmlDoc, _ := content.ReadFile("src/index.html")
	browser := lib.NewBrowser(string(htmlDoc), settings)

	// Define browser WebView script bindings.
	browser.BindFuncVoid("browser_Load", func(arg ...string) {
		fileData, err := content.ReadFile("src/" + lib.StripSlash(arg[0]))

		if err == nil {
			browser.LoadHtml(string(fileData))
		}
	})

	browser.LoadScript(lib.InitScript())
	browser.Open()
}
