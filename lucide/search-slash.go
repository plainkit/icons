package lucide

import (
	html "github.com/plainkit/html"
)

// SearchSlash creates a Search Slash Lucide icon.
func SearchSlash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-search-slash", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m13.5 8.5-5 5"))),
		html.Child(html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("8"))),
		html.Child(html.SvgPath(html.AD("m21 21-4.3-4.3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
