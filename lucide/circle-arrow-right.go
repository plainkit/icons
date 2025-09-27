package lucide

import (
	html "github.com/plainkit/html"
)

// CircleArrowRight creates a Circle Arrow Right Lucide icon.
func CircleArrowRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-arrow-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("m12 16 4-4-4-4"))),
		html.Child(html.SvgPath(html.AD("M8 12h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
