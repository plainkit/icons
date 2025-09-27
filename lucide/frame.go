package lucide

import (
	html "github.com/plainkit/html"
)

// Frame creates a Frame Lucide icon.
func Frame(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-frame", args)
	children := []html.SvgArg{
		html.Child(html.SvgLine(html.AX1("22"), html.AX2("2"), html.AY1("6"), html.AY2("6"))),
		html.Child(html.SvgLine(html.AX1("22"), html.AX2("2"), html.AY1("18"), html.AY2("18"))),
		html.Child(html.SvgLine(html.AX1("6"), html.AX2("6"), html.AY1("2"), html.AY2("22"))),
		html.Child(html.SvgLine(html.AX1("18"), html.AX2("18"), html.AY1("2"), html.AY2("22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
