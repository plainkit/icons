package lucide

import (
	html "github.com/plainkit/html"
)

// Dribbble creates a Dribbble Lucide icon.
func Dribbble(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dribbble", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M19.13 5.09C15.22 9.14 10 10.44 2.25 10.94")),
		html.SvgPath(html.AD("M21.75 12.84c-6.62-1.41-12.14 1-16.38 6.32")),
		html.SvgPath(html.AD("M8.56 2.75c4.37 6 6 9.42 8 17.72")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
