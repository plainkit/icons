package lucide

import (
	html "github.com/plainkit/html"
)

// Infinity creates a Infinity Lucide icon.
func Infinity(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-infinity", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 16c5 0 7-8 12-8a4 4 0 0 1 0 8c-5 0-7-8-12-8a4 4 0 1 0 0 8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
