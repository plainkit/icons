package lucide

import (
	html "github.com/plainkit/html"
)

// Option creates a Option Lucide icon.
func Option(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-option", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 3h6l6 18h6"))),
		html.Child(html.SvgPath(html.AD("M14 3h7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
