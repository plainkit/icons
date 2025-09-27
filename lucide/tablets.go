package lucide

import (
	html "github.com/plainkit/html"
)

// Tablets creates a Tablets Lucide icon.
func Tablets(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tablets", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("7"), html.ACy("7"), html.AR("5"))),
		html.Child(html.SvgCircle(html.ACx("17"), html.ACy("17"), html.AR("5"))),
		html.Child(html.SvgPath(html.AD("M12 17h10"))),
		html.Child(html.SvgPath(html.AD("m3.46 10.54 7.08-7.08"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
