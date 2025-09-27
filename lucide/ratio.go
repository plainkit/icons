package lucide

import (
	html "github.com/plainkit/html"
)

// Ratio creates a Ratio Lucide icon.
func Ratio(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ratio", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("12"), html.AHeight("20"), html.AX("6"), html.AY("2"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
