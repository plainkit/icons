package lucide

import (
	html "github.com/plainkit/html"
)

// SquareStop creates a Square Stop Lucide icon.
func SquareStop(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-stop", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgRect(html.AWidth("6"), html.AHeight("6"), html.AX("9"), html.AY("9"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
