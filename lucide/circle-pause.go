package lucide

import (
	html "github.com/plainkit/html"
)

// CirclePause creates a Circle Pause Lucide icon.
func CirclePause(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-pause", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgLine(html.AX1("10"), html.AX2("10"), html.AY1("15"), html.AY2("9")),
		html.SvgLine(html.AX1("14"), html.AX2("14"), html.AY1("15"), html.AY2("9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
