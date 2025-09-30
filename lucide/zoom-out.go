package lucide

import (
	html "github.com/plainkit/html"
)

// ZoomOut creates a Zoom Out Lucide icon.
func ZoomOut(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-zoom-out", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("8")),
		html.SvgLine(html.AX1("21"), html.AX2("16.65"), html.AY1("21"), html.AY2("16.65")),
		html.SvgLine(html.AX1("8"), html.AX2("14"), html.AY1("11"), html.AY2("11")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
