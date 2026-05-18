# WebView Source Builder

Create self-contained binary of [HTML](https://en.wikipedia.org/wiki/HTML) sources in a [WebView](https://en.wikipedia.org/wiki/WebView) :warning: Work In Progress :warning:

## Dependencies

- [Go](https://golang.org)

### Debian/Ubuntu

The following dependencies are required in order to build for Debian-based operating systems.  For alternate OS's (e.g. BSD, Windows) refer to the [webview preqequisites](https://github.com/webview/webview?tab=readme-ov-file#prerequisites) install instructions.

    $ apt-get install -y libgtk-3-dev libwebkit2gtk-4.0-dev

## Supported elements

Currently the only HTML element not supported is inline `<style>` usage due to browser [Cross-origin resource sharing](https://www.w3.org/TR/2020/SPSD-cors-20200602) (CORS) restrictions.

As an alternative you should include your stylesheets using:

```txt
<link rel="stylesheet" type="text/css" href="path/to/file.css" />
```

## Quick and Easy

Compile the Go application and run the example in one command:

    $ make

## Build from source

Install the new build using [gmake](https://www.gnu.org/software/make).

    $ make install

Cross-compile to support [Windows](https://golang.org/dl/go1.15.6.windows-amd64.msi), [OSX](https://golang.org/dl/go1.15.6.darwin-amd64.pkg), [etc](https://golang.org/dl) ..

    $ make build-<darwin|linux|windows>

## Running the application

Once compiled it should be as easy as..

    $ webview-src

Note: Using `--debug` will enable WebView browser [Developer Tools](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Tools_and_setup/What_are_browser_developer_tools).

## Go application structure

```text
lib    // Go package dependencies.
src    // HTML markup and related sources.
src.go // main
```

## References

- [webview_go](https://github.com/webview/webview_go) - Go language binding for the webview library.

## Contributions

If you fix a bug, or have a code you want to contribute, please send a pull-request with your changes. (Note: Before committing your code please ensure that you run [golint](https://github.com/golang/lint) and [gofmt](https://pkg.go.dev/cmd/gofmt) on contributed files).

## Versioning

This package is maintained under the [Semantic Versioning](https://semver.org) guidelines.

## License and Warranty

This package is distributed in the hope that it will be useful, but without any warranty; without even the implied warranty of merchantability or fitness for a particular purpose.

_go-webview-src-builder_ is provided under the terms of the [MIT license](http://www.opensource.org/licenses/mit-license.php)

## Author

[Marc S. Brooks](https://github.com/nuxy)
