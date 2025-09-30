package lucide

import (
	html "github.com/plainkit/html"
)

// TreeDeciduous creates a Tree Deciduous Lucide icon.
func TreeDeciduous(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tree-deciduous", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 19a4 4 0 0 1-2.24-7.32A3.5 3.5 0 0 1 9 6.03V6a3 3 0 1 1 6 0v.04a3.5 3.5 0 0 1 3.24 5.65A4 4 0 0 1 16 19Z")),
		html.SvgPath(html.AD("M12 19v3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
