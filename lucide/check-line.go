package lucide

import (
	html "github.com/plainkit/html"
)

// CheckLine creates a Check Line Lucide icon.
func CheckLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-check-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 4L9 15"))),
		html.Child(html.SvgPath(html.AD("M21 19L3 19"))),
		html.Child(html.SvgPath(html.AD("M9 15L4 10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
