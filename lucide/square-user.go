package lucide

import (
	html "github.com/plainkit/html"
)

// SquareUser creates a Square User Lucide icon.
func SquareUser(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-user", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M7 21v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
