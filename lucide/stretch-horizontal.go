package lucide

import (
	html "github.com/plainkit/html"
)

// StretchHorizontal creates a Stretch Horizontal Lucide icon.
func StretchHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-stretch-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("6"), html.AX("2"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("6"), html.AX("2"), html.AY("14"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
