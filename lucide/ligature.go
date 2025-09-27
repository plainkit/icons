package lucide

import (
	html "github.com/plainkit/html"
)

// Ligature creates a Ligature Lucide icon.
func Ligature(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ligature", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 12h2v8"))),
		html.Child(html.SvgPath(html.AD("M14 20h4"))),
		html.Child(html.SvgPath(html.AD("M6 12h4"))),
		html.Child(html.SvgPath(html.AD("M6 20h4"))),
		html.Child(html.SvgPath(html.AD("M8 20V8a4 4 0 0 1 7.464-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
