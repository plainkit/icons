package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowsUpFromLine creates a Arrows Up From Line Lucide icon.
func ArrowsUpFromLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrows-up-from-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m4 6 3-3 3 3"))),
		html.Child(html.SvgPath(html.AD("M7 17V3"))),
		html.Child(html.SvgPath(html.AD("m14 6 3-3 3 3"))),
		html.Child(html.SvgPath(html.AD("M17 17V3"))),
		html.Child(html.SvgPath(html.AD("M4 21h16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
