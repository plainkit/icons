package lucide

import (
	html "github.com/plainkit/html"
)

// CircleParking creates a Circle Parking Lucide icon.
func CircleParking(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-parking", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M9 17V7h4a3 3 0 0 1 0 6H9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
