package lucide

import (
	html "github.com/plainkit/html"
)

// WindArrowDown creates a Wind Arrow Down Lucide icon.
func WindArrowDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wind-arrow-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 2v8")),
		html.SvgPath(html.AD("M12.8 21.6A2 2 0 1 0 14 18H2")),
		html.SvgPath(html.AD("M17.5 10a2.5 2.5 0 1 1 2 4H2")),
		html.SvgPath(html.AD("m6 6 4 4 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
