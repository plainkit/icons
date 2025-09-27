package lucide

import (
	html "github.com/plainkit/html"
)

// SquareParkingOff creates a Square Parking Off Lucide icon.
func SquareParkingOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-parking-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.6 3.6A2 2 0 0 1 5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-.59 1.41"))),
		html.Child(html.SvgPath(html.AD("M3 8.7V19a2 2 0 0 0 2 2h10.3"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M13 13a3 3 0 1 0 0-6H9v2"))),
		html.Child(html.SvgPath(html.AD("M9 17v-2.3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
