package lucide

import (
	html "github.com/plainkit/html"
)

// Donut creates a Donut Lucide icon.
func Donut(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-donut", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20.5 10a2.5 2.5 0 0 1-2.4-3H18a2.95 2.95 0 0 1-2.6-4.4 10 10 0 1 0 6.3 7.1c-.3.2-.8.3-1.2.3"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
