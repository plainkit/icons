package lucide

import (
	html "github.com/plainkit/html"
)

// UserMinus creates a User Minus Lucide icon.
func UserMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-minus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2")),
		html.SvgCircle(html.ACx("9"), html.ACy("7"), html.AR("4")),
		html.SvgLine(html.AX1("22"), html.AX2("16"), html.AY1("11"), html.AY2("11")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
