package lucide

import (
	html "github.com/plainkit/html"
)

// SunSnow creates a Sun Snow Lucide icon.
func SunSnow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sun-snow", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 21v-1"))),
		html.Child(html.SvgPath(html.AD("M10 4V3"))),
		html.Child(html.SvgPath(html.AD("M10 9a3 3 0 0 0 0 6"))),
		html.Child(html.SvgPath(html.AD("m14 20 1.25-2.5L18 18"))),
		html.Child(html.SvgPath(html.AD("m14 4 1.25 2.5L18 6"))),
		html.Child(html.SvgPath(html.AD("m17 21-3-6 1.5-3H22"))),
		html.Child(html.SvgPath(html.AD("m17 3-3 6 1.5 3"))),
		html.Child(html.SvgPath(html.AD("M2 12h1"))),
		html.Child(html.SvgPath(html.AD("m20 10-1.5 2 1.5 2"))),
		html.Child(html.SvgPath(html.AD("m3.64 18.36.7-.7"))),
		html.Child(html.SvgPath(html.AD("m4.34 6.34-.7-.7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
