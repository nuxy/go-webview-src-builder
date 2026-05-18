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
  updateCSS();
  updateImages();
  updateLinks();
});

/**
 * Replace LINK stylesheets with data: equivalent.
 */
async function updateCSS() {
  var nodes = document.querySelectorAll('link[href]');

  for (var i = 0; i < nodes.length; i++) {
    var node = nodes[i];
    var data = await window.browser_Get(node.getAttribute('href'));

    node.setAttribute('href', data);
  }
}

/**
 * Replace HTML images with data: equivalent.
 */
async function updateImages() {
  var nodes = document.querySelectorAll('img[src]');

  for (var i = 0; i < nodes.length; i++) {
    var node = nodes[i];
    var data = await window.browser_Get(node.getAttribute('src'));

    node.setAttribute('src', data);
  }
}

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
