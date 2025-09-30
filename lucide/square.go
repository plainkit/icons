package lucide

import (
	html "github.com/plainkit/html"
)

// Square creates a Square Lucide icon.
func Square(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
