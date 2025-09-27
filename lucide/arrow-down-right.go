package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownRight creates a Arrow Down Right Lucide icon.
func ArrowDownRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m7 7 10 10"))),
		html.Child(html.SvgPath(html.AD("M17 7v10H7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
