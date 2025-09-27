package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDown creates a Arrow Down Lucide icon.
func ArrowDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 5v14"))),
		html.Child(html.SvgPath(html.AD("m19 12-7 7-7-7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
