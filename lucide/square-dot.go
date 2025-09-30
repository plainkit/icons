package lucide

import (
	html "github.com/plainkit/html"
)

// SquareDot creates a Square Dot Lucide icon.
func SquareDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-dot", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
