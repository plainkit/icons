package lucide

import (
	html "github.com/plainkit/html"
)

// Grid2x2X creates a Grid 2x2 X Lucide icon.
func Grid2x2X(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grid-2x2-x", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 3v17a1 1 0 0 1-1 1H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6a1 1 0 0 1-1 1H3"))),
		html.Child(html.SvgPath(html.AD("m16 16 5 5"))),
		html.Child(html.SvgPath(html.AD("m16 21 5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
