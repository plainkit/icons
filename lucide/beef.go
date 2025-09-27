package lucide

import (
	html "github.com/plainkit/html"
)

// Beef creates a Beef Lucide icon.
func Beef(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-beef", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16.4 13.7A6.5 6.5 0 1 0 6.28 6.6c-1.1 3.13-.78 3.9-3.18 6.08A3 3 0 0 0 5 18c4 0 8.4-1.8 11.4-4.3"))),
		html.Child(html.SvgPath(html.AD("m18.5 6 2.19 4.5a6.48 6.48 0 0 1-2.29 7.2C15.4 20.2 11 22 7 22a3 3 0 0 1-2.68-1.66L2.4 16.5"))),
		html.Child(html.SvgCircle(html.ACx("12.5"), html.ACy("8.5"), html.AR("2.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
