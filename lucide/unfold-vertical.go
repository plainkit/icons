package lucide

import (
	html "github.com/plainkit/html"
)

// UnfoldVertical creates a Unfold Vertical Lucide icon.
func UnfoldVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-unfold-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 22v-6"))),
		html.Child(html.SvgPath(html.AD("M12 8V2"))),
		html.Child(html.SvgPath(html.AD("M4 12H2"))),
		html.Child(html.SvgPath(html.AD("M10 12H8"))),
		html.Child(html.SvgPath(html.AD("M16 12h-2"))),
		html.Child(html.SvgPath(html.AD("M22 12h-2"))),
		html.Child(html.SvgPath(html.AD("m15 19-3 3-3-3"))),
		html.Child(html.SvgPath(html.AD("m15 5-3-3-3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
