package lucide

import (
	html "github.com/plainkit/html"
)

// SearchX creates a Search X Lucide icon.
func SearchX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-search-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m13.5 8.5-5 5")),
		html.SvgPath(html.AD("m8.5 8.5 5 5")),
		html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("8")),
		html.SvgPath(html.AD("m21 21-4.3-4.3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
