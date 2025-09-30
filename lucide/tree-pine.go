package lucide

import (
	html "github.com/plainkit/html"
)

// TreePine creates a Tree Pine Lucide icon.
func TreePine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tree-pine", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m17 14 3 3.3a1 1 0 0 1-.7 1.7H4.7a1 1 0 0 1-.7-1.7L7 14h-.3a1 1 0 0 1-.7-1.7L9 9h-.2A1 1 0 0 1 8 7.3L12 3l4 4.3a1 1 0 0 1-.8 1.7H15l3 3.3a1 1 0 0 1-.7 1.7H17Z")),
		html.SvgPath(html.AD("M12 22v-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
