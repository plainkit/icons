package lucide

import (
	html "github.com/plainkit/html"
)

// Heading3 creates a Heading 3 Lucide icon.
func Heading3(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heading-3", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 12h8")),
		html.SvgPath(html.AD("M4 18V6")),
		html.SvgPath(html.AD("M12 18V6")),
		html.SvgPath(html.AD("M17.5 10.5c1.7-1 3.5 0 3.5 1.5a2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("M17 17.5c2 1.5 4 .3 4-1.5a2 2 0 0 0-2-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
