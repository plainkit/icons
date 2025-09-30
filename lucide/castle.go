package lucide

import (
	html "github.com/plainkit/html"
)

// Castle creates a Castle Lucide icon.
func Castle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-castle", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 5V3")),
		html.SvgPath(html.AD("M14 5V3")),
		html.SvgPath(html.AD("M15 21v-3a3 3 0 0 0-6 0v3")),
		html.SvgPath(html.AD("M18 3v8")),
		html.SvgPath(html.AD("M18 5H6")),
		html.SvgPath(html.AD("M22 11H2")),
		html.SvgPath(html.AD("M22 9v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9")),
		html.SvgPath(html.AD("M6 3v8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
