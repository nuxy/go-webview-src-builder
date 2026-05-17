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
window.addEventListener('load', function() {
  var links = document.querySelectorAll('a[href]');

  for (var i = 0; i < links.length; i++) {
    var link = links[i];

    link.addEventListener('click', function(event) {
      event.preventDefault();
	
      window.browser_Load(this.getAttribute('href'));
    });
  }
});
`
}
