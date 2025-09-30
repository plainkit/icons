package lucide

import (
	html "github.com/plainkit/html"
)

// Smile creates a Smile Lucide icon.
func Smile(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-smile", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M8 14s1.5 2 4 2 4-2 4-2")),
		html.SvgLine(html.AX1("9"), html.AX2("9.01"), html.AY1("9"), html.AY2("9")),
		html.SvgLine(html.AX1("15"), html.AX2("15.01"), html.AY1("9"), html.AY2("9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
