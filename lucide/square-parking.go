package lucide

import (
	html "github.com/plainkit/html"
)

// SquareParking creates a Square Parking Lucide icon.
func SquareParking(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-parking", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M9 17V7h4a3 3 0 0 1 0 6H9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
