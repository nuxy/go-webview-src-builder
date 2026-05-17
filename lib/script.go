//
//  go-webview-src-builder
//  Create a self-contained binary of HTML sources in a WebView
//
//  Copyright 2026, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package lib

// Returns WebView injection JavaScript.
func InitScript() string {
	return `

/**
 * Update page content using Go WebView bindings.
 */
window.addEventListener('load', function() {
  updateLinks();
});

/**
 * Replace HTML anchor link locations.
 */
function updateLinks() {
  var nodes = document.querySelectorAll('a[href]');

  for (var i = 0; i < nodes.length; i++) {
    var node = nodes[i];

    node.addEventListener('click', async function(event) {
      event.preventDefault();

      await window.browser_Load(this.getAttribute('href'));
    });
  }
}

`
}
