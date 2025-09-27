package lucide

import (
	html "github.com/plainkit/html"
)

// Caravan creates a Caravan Lucide icon.
func Caravan(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-caravan", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 19V9a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v8a2 2 0 0 0 2 2h2"))),
		html.Child(html.SvgPath(html.AD("M2 9h3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H2"))),
		html.Child(html.SvgPath(html.AD("M22 17v1a1 1 0 0 1-1 1H10v-9a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v9"))),
		html.Child(html.SvgCircle(html.ACx("8"), html.ACy("19"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
