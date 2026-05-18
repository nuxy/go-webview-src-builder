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
	"regexp"
	"strconv"
)

// Converts a valid string value to boolean.
func StrToBool(v string) bool {
	_v, err := strconv.ParseBool(v)

	if err != nil {
		log.Fatal(err)
	}

	return _v
}

// Removes .. and / prefix from a string value.
func CleanPath(v string) string {
	re := regexp.MustCompile(`^(\.\.|\/)+`)
	return re.ReplaceAllString(v, "")
}
