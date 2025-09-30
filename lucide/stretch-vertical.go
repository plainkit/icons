package lucide

import (
	html "github.com/plainkit/html"
)

// StretchVertical creates a Stretch Vertical Lucide icon.
func StretchVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-stretch-vertical", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("6"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2")),
		html.SvgRect(html.AWidth("6"), html.AHeight("20"), html.AX("14"), html.AY("2"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
