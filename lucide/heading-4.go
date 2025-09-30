package lucide

import (
	html "github.com/plainkit/html"
)

// Heading4 creates a Heading 4 Lucide icon.
func Heading4(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heading-4", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 18V6")),
		html.SvgPath(html.AD("M17 10v3a1 1 0 0 0 1 1h3")),
		html.SvgPath(html.AD("M21 10v8")),
		html.SvgPath(html.AD("M4 12h8")),
		html.SvgPath(html.AD("M4 18V6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
