package lucide

import (
	html "github.com/plainkit/html"
)

// UserX creates a User X Lucide icon.
func UserX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2")),
		html.SvgCircle(html.ACx("9"), html.ACy("7"), html.AR("4")),
		html.SvgLine(html.AX1("17"), html.AX2("22"), html.AY1("8"), html.AY2("13")),
		html.SvgLine(html.AX1("22"), html.AX2("17"), html.AY1("8"), html.AY2("13")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
