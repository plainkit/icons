package lucide

import (
	html "github.com/plainkit/html"
)

// Grid2x2Plus creates a Grid 2x2 Plus Lucide icon.
func Grid2x2Plus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grid-2x2-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 3v17a1 1 0 0 1-1 1H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6a1 1 0 0 1-1 1H3"))),
		html.Child(html.SvgPath(html.AD("M16 19h6"))),
		html.Child(html.SvgPath(html.AD("M19 22v-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
