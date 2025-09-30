package lucide

import (
	html "github.com/plainkit/html"
)

// FoldVertical creates a Fold Vertical Lucide icon.
func FoldVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fold-vertical", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 22v-6")),
		html.SvgPath(html.AD("M12 8V2")),
		html.SvgPath(html.AD("M4 12H2")),
		html.SvgPath(html.AD("M10 12H8")),
		html.SvgPath(html.AD("M16 12h-2")),
		html.SvgPath(html.AD("M22 12h-2")),
		html.SvgPath(html.AD("m15 19-3-3-3 3")),
		html.SvgPath(html.AD("m15 5-3 3-3-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
