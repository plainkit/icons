package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowRight creates a Arrow Right Lucide icon.
func ArrowRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 12h14"))),
		html.Child(html.SvgPath(html.AD("m12 5 7 7-7 7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
