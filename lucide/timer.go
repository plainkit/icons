package lucide

import (
	html "github.com/plainkit/html"
)

// Timer creates a Timer Lucide icon.
func Timer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-timer", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("10"), html.AX2("14"), html.AY1("2"), html.AY2("2")),
		html.SvgLine(html.AX1("12"), html.AX2("15"), html.AY1("14"), html.AY2("11")),
		html.SvgCircle(html.ACx("12"), html.ACy("14"), html.AR("8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
