package lucide

import (
	html "github.com/plainkit/html"
)

// Tornado creates a Tornado Lucide icon.
func Tornado(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tornado", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 4H3"))),
		html.Child(html.SvgPath(html.AD("M18 8H6"))),
		html.Child(html.SvgPath(html.AD("M19 12H9"))),
		html.Child(html.SvgPath(html.AD("M16 16h-6"))),
		html.Child(html.SvgPath(html.AD("M11 20H9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
