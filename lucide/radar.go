package lucide

import (
	html "github.com/plainkit/html"
)

// Radar creates a Radar Lucide icon.
func Radar(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-radar", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M19.07 4.93A10 10 0 0 0 6.99 3.34"))),
		html.Child(html.SvgPath(html.AD("M4 6h.01"))),
		html.Child(html.SvgPath(html.AD("M2.29 9.62A10 10 0 1 0 21.31 8.35"))),
		html.Child(html.SvgPath(html.AD("M16.24 7.76A6 6 0 1 0 8.23 16.67"))),
		html.Child(html.SvgPath(html.AD("M12 18h.01"))),
		html.Child(html.SvgPath(html.AD("M17.99 11.66A6 6 0 0 1 15.77 16.67"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("m13.41 10.59 5.66-5.66"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
