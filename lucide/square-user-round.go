package lucide

import (
	html "github.com/plainkit/html"
)

// SquareUserRound creates a Square User Round Lucide icon.
func SquareUserRound(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-user-round", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 21a6 6 0 0 0-12 0"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("11"), html.AR("4"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
