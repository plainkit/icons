package lucide

import (
	html "github.com/plainkit/html"
)

// SlidersVertical creates a Sliders Vertical Lucide icon.
func SlidersVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sliders-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 8h4"))),
		html.Child(html.SvgPath(html.AD("M12 21v-9"))),
		html.Child(html.SvgPath(html.AD("M12 8V3"))),
		html.Child(html.SvgPath(html.AD("M17 16h4"))),
		html.Child(html.SvgPath(html.AD("M19 12V3"))),
		html.Child(html.SvgPath(html.AD("M19 21v-5"))),
		html.Child(html.SvgPath(html.AD("M3 14h4"))),
		html.Child(html.SvgPath(html.AD("M5 10V3"))),
		html.Child(html.SvgPath(html.AD("M5 21v-7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
