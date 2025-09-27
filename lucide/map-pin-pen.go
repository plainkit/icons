package lucide

import (
	html "github.com/plainkit/html"
)

// MapPinPen creates a Map Pin Pen Lucide icon.
func MapPinPen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-pin-pen", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17.97 9.304A8 8 0 0 0 2 10c0 4.69 4.887 9.562 7.022 11.468"))),
		html.Child(html.SvgPath(html.AD("M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("10"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
