package lucide

import (
	html "github.com/plainkit/html"
)

// ZoomIn creates a Zoom In Lucide icon.
func ZoomIn(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-zoom-in", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("8"))),
		html.Child(html.SvgLine(html.AX1("21"), html.AX2("16.65"), html.AY1("21"), html.AY2("16.65"))),
		html.Child(html.SvgLine(html.AX1("11"), html.AX2("11"), html.AY1("8"), html.AY2("14"))),
		html.Child(html.SvgLine(html.AX1("8"), html.AX2("14"), html.AY1("11"), html.AY2("11"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
