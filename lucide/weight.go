package lucide

import (
	html "github.com/plainkit/html"
)

// Weight creates a Weight Lucide icon.
func Weight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-weight", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("5"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M6.5 8a2 2 0 0 0-1.905 1.46L2.1 18.5A2 2 0 0 0 4 21h16a2 2 0 0 0 1.925-2.54L19.4 9.5A2 2 0 0 0 17.48 8Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
