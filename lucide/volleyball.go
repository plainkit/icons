package lucide

import (
	html "github.com/plainkit/html"
)

// Volleyball creates a Volleyball Lucide icon.
func Volleyball(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-volleyball", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11.1 7.1a16.55 16.55 0 0 1 10.9 4"))),
		html.Child(html.SvgPath(html.AD("M12 12a12.6 12.6 0 0 1-8.7 5"))),
		html.Child(html.SvgPath(html.AD("M16.8 13.6a16.55 16.55 0 0 1-9 7.5"))),
		html.Child(html.SvgPath(html.AD("M20.7 17a12.8 12.8 0 0 0-8.7-5 13.3 13.3 0 0 1 0-10"))),
		html.Child(html.SvgPath(html.AD("M6.3 3.8a16.55 16.55 0 0 0 1.9 11.5"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
