package lucide

import (
	html "github.com/plainkit/html"
)

// SquareStack creates a Square Stack Lucide icon.
func SquareStack(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-stack", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 10c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2")),
		html.SvgPath(html.AD("M10 16c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("14"), html.AY("14"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
