package lucide

import (
	html "github.com/plainkit/html"
)

// Wine creates a Wine Lucide icon.
func Wine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wine", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 22h8")),
		html.SvgPath(html.AD("M7 10h10")),
		html.SvgPath(html.AD("M12 15v7")),
		html.SvgPath(html.AD("M12 15a5 5 0 0 0 5-5c0-2-.5-4-2-8H9c-1.5 4-2 6-2 8a5 5 0 0 0 5 5Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
