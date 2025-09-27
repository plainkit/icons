package lucide

import (
	html "github.com/plainkit/html"
)

// CircleDollarSign creates a Circle Dollar Sign Lucide icon.
func CircleDollarSign(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-dollar-sign", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8"))),
		html.Child(html.SvgPath(html.AD("M12 18V6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
