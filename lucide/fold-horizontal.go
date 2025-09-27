package lucide

import (
	html "github.com/plainkit/html"
)

// FoldHorizontal creates a Fold Horizontal Lucide icon.
func FoldHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-fold-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 12h6"))),
		html.Child(html.SvgPath(html.AD("M22 12h-6"))),
		html.Child(html.SvgPath(html.AD("M12 2v2"))),
		html.Child(html.SvgPath(html.AD("M12 8v2"))),
		html.Child(html.SvgPath(html.AD("M12 14v2"))),
		html.Child(html.SvgPath(html.AD("M12 20v2"))),
		html.Child(html.SvgPath(html.AD("m19 9-3 3 3 3"))),
		html.Child(html.SvgPath(html.AD("m5 15 3-3-3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
