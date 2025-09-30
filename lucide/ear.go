package lucide

import (
	html "github.com/plainkit/html"
)

// Ear creates a Ear Lucide icon.
func Ear(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ear", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 8.5a6.5 6.5 0 1 1 13 0c0 6-6 6-6 10a3.5 3.5 0 1 1-7 0")),
		html.SvgPath(html.AD("M15 8.5a2.5 2.5 0 0 0-5 0v1a2 2 0 1 1 0 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
