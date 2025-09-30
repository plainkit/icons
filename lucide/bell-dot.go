package lucide

import (
	html "github.com/plainkit/html"
)

// BellDot creates a Bell Dot Lucide icon.
func BellDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bell-dot", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.268 21a2 2 0 0 0 3.464 0")),
		html.SvgPath(html.AD("M13.916 2.314A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.74 7.327A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673 9 9 0 0 1-.585-.665")),
		html.SvgCircle(html.ACx("18"), html.ACy("8"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
