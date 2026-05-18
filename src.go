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
	"path"
	"regexp"
	"strings"

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
	htmlDoc, _ := readFile("index.html")
	browser := lib.NewBrowser(string(htmlDoc), settings)

	// Define browser WebView script bindings.
	browser.BindFuncReturn("browser_Get", func(arg ...string) string {
		fileData, err := readFile(arg[0])

		if err == nil {
			fileType := path.Ext(arg[0])

			if fileType == ".css" {
				fileData = []byte(parseUrlFunc(string(fileData)))
			}

			return lib.EncodeData(fileData, fileType)
		}

		return arg[0] // no changes
	})

	browser.BindFuncVoid("browser_Load", func(arg ...string) {
		fileData, err := readFile(arg[0])

		if err == nil {
			browser.LoadHtml(string(fileData))
		}
	})

	browser.LoadScript(lib.InitScript())
	browser.Open()
}

// Get embedded content for a given path.
func readFile(v string) ([]byte, error) {
	return content.ReadFile("src/" + lib.CleanPath(v))
}

// Replace CSS url() with HTTP data:
func parseUrlFunc(v string) []byte {
	re := regexp.MustCompile(`url\(['"]?([^}]*)['"]\)`)
	matches := re.FindAllStringSubmatch(v, -1)

	for i := 0; i < len(matches); i++ {
		filePath := matches[i][1]
		fileData, err := readFile(filePath)

		if err == nil {
			fileType := path.Ext(filePath)

			encodedData := lib.EncodeData(fileData, fileType)

			v = strings.Replace(v, filePath, encodedData, 1)
		}
	}

	return []byte(v)
}
