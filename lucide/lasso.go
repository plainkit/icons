package lucide

import (
	html "github.com/plainkit/html"
)

// Lasso creates a Lasso Lucide icon.
func Lasso(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lasso", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.704 14.467A10 8 0 0 1 2 10a10 8 0 0 1 20 0 10 8 0 0 1-10 8 10 8 0 0 1-5.181-1.158"))),
		html.Child(html.SvgPath(html.AD("M7 22a5 5 0 0 1-2-3.994"))),
		html.Child(html.SvgCircle(html.ACx("5"), html.ACy("16"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
