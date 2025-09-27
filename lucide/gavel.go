package lucide

import (
	html "github.com/plainkit/html"
)

// Gavel creates a Gavel Lucide icon.
func Gavel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gavel", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m14 13-8.381 8.38a1 1 0 0 1-3.001-3l8.384-8.381"))),
		html.Child(html.SvgPath(html.AD("m16 16 6-6"))),
		html.Child(html.SvgPath(html.AD("m21.5 10.5-8-8"))),
		html.Child(html.SvgPath(html.AD("m8 8 6-6"))),
		html.Child(html.SvgPath(html.AD("m8.5 7.5 8 8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
