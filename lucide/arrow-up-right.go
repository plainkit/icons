package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpRight creates a Arrow Up Right Lucide icon.
func ArrowUpRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 7h10v10"))),
		html.Child(html.SvgPath(html.AD("M7 17 17 7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
