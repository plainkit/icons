package lucide

import (
	html "github.com/plainkit/html"
)

// SplinePointer creates a Spline Pointer Lucide icon.
func SplinePointer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-spline-pointer", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12.034 12.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.943l-3.444 1.068a1 1 0 0 0-.66.66l-1.067 3.443a.5.5 0 0 1-.943.033z")),
		html.SvgPath(html.AD("M5 17A12 12 0 0 1 17 5")),
		html.SvgCircle(html.ACx("19"), html.ACy("5"), html.AR("2")),
		html.SvgCircle(html.ACx("5"), html.ACy("19"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
