package lucide

import (
	html "github.com/plainkit/html"
)

// SearchCheck creates a Search Check Lucide icon.
func SearchCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-search-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m8 11 2 2 4-4"))),
		html.Child(html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("8"))),
		html.Child(html.SvgPath(html.AD("m21 21-4.3-4.3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
