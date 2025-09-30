package lucide

import (
	html "github.com/plainkit/html"
)

// SearchCode creates a Search Code Lucide icon.
func SearchCode(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-search-code", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m13 13.5 2-2.5-2-2.5")),
		html.SvgPath(html.AD("m21 21-4.3-4.3")),
		html.SvgPath(html.AD("M9 8.5 7 11l2 2.5")),
		html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
