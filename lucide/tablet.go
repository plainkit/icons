package lucide

import (
	html "github.com/plainkit/html"
)

// Tablet creates a Tablet Lucide icon.
func Tablet(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tablet", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12.01"), html.AY1("18"), html.AY2("18"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
