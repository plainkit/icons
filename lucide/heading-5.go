package lucide

import (
	html "github.com/plainkit/html"
)

// Heading5 creates a Heading 5 Lucide icon.
func Heading5(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heading-5", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 12h8")),
		html.SvgPath(html.AD("M4 18V6")),
		html.SvgPath(html.AD("M12 18V6")),
		html.SvgPath(html.AD("M17 13v-3h4")),
		html.SvgPath(html.AD("M17 17.7c.4.2.8.3 1.3.3 1.5 0 2.7-1.1 2.7-2.5S19.8 13 18.3 13H17")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
