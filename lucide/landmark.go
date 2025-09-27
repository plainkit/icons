package lucide

import (
	html "github.com/plainkit/html"
)

// Landmark creates a Landmark Lucide icon.
func Landmark(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-landmark", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 18v-7"))),
		html.Child(html.SvgPath(html.AD("M11.12 2.198a2 2 0 0 1 1.76.006l7.866 3.847c.476.233.31.949-.22.949H3.474c-.53 0-.695-.716-.22-.949z"))),
		html.Child(html.SvgPath(html.AD("M14 18v-7"))),
		html.Child(html.SvgPath(html.AD("M18 18v-7"))),
		html.Child(html.SvgPath(html.AD("M3 22h18"))),
		html.Child(html.SvgPath(html.AD("M6 18v-7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
