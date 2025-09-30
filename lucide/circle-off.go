package lucide

import (
	html "github.com/plainkit/html"
)

// CircleOff creates a Circle Off Lucide icon.
func CircleOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M8.35 2.69A10 10 0 0 1 21.3 15.65")),
		html.SvgPath(html.AD("M19.08 19.08A10 10 0 1 1 4.92 4.92")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
