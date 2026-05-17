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
	"encoding/base64"
	"mime"
	"net/http"
)

// Generate browser HTTP data: URL
func EncodeData(v []byte, fileExt string) string {
	var mimeType string = getMimeType(fileExt)

	if mimeType == "" {
		mimeType = http.DetectContentType(v)
	}

	return "data:" + mimeType + ";base64," + toBase64(v)
}

// Get MIME type for given file type.
func getMimeType(v string) string {
	return mime.TypeByExtension(v)
}

// Convert bytes to Base64 string value.
func toBase64(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}
