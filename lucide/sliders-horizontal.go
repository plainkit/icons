package lucide

import (
	html "github.com/plainkit/html"
)

// SlidersHorizontal creates a Sliders Horizontal Lucide icon.
func SlidersHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sliders-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 5H3"))),
		html.Child(html.SvgPath(html.AD("M12 19H3"))),
		html.Child(html.SvgPath(html.AD("M14 3v4"))),
		html.Child(html.SvgPath(html.AD("M16 17v4"))),
		html.Child(html.SvgPath(html.AD("M21 12h-9"))),
		html.Child(html.SvgPath(html.AD("M21 19h-5"))),
		html.Child(html.SvgPath(html.AD("M21 5h-7"))),
		html.Child(html.SvgPath(html.AD("M8 10v4"))),
		html.Child(html.SvgPath(html.AD("M8 12H3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
