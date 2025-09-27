package lucide

import (
	html "github.com/plainkit/html"
)

// UserPlus creates a User Plus Lucide icon.
func UserPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"))),
		html.Child(html.SvgCircle(html.ACx("9"), html.ACy("7"), html.AR("4"))),
		html.Child(html.SvgLine(html.AX1("19"), html.AX2("19"), html.AY1("8"), html.AY2("14"))),
		html.Child(html.SvgLine(html.AX1("22"), html.AX2("16"), html.AY1("11"), html.AY2("11"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
