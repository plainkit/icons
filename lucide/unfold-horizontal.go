package lucide

import (
	html "github.com/plainkit/html"
)

// UnfoldHorizontal creates a Unfold Horizontal Lucide icon.
func UnfoldHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-unfold-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 12h6"))),
		html.Child(html.SvgPath(html.AD("M8 12H2"))),
		html.Child(html.SvgPath(html.AD("M12 2v2"))),
		html.Child(html.SvgPath(html.AD("M12 8v2"))),
		html.Child(html.SvgPath(html.AD("M12 14v2"))),
		html.Child(html.SvgPath(html.AD("M12 20v2"))),
		html.Child(html.SvgPath(html.AD("m19 15 3-3-3-3"))),
		html.Child(html.SvgPath(html.AD("m5 9-3 3 3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
