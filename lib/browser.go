//
//  go-webview-src-builder
//  Create a self-contained binary of HTML sources in a WebView
//
//  Copyright 2026, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package lib

import (
	"log"
	"strings"

	webview "github.com/webview/webview_go"
)

// Browser declared data types.
type BrowserSettings struct {
	Title  string
	Height int
	Width  int
	Resize bool
    Debug  bool
}

type Browser struct {
	WebView  webview.WebView
	document string
	settings BrowserSettings
}

// NewBrowser creates a WebView instance.
func NewBrowser(htmlMarkup string, settings BrowserSettings) *Browser {
	browser := &Browser{settings: settings}
	browser.document = htmlMarkup
	browser.init()
	return browser
}

// Initialize a new WebView window.
func (browser *Browser) init() {
	var webviewHint webview.Hint = webview.HintFixed

	if browser.settings.Resize {
		webviewHint = webview.HintNone
	}

	browser.WebView = webview.New(browser.settings.Debug)
	browser.WebView.SetTitle(browser.settings.Title)
	browser.WebView.SetSize(browser.settings.Width, browser.settings.Height, webviewHint)
	browser.WebView.SetHtml(browser.document)
}

// Launch the WebView window.
func (browser *Browser) Open() {
	browser.WebView.Run()
}
