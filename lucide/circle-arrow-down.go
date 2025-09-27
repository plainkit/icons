package lucide

import (
	html "github.com/plainkit/html"
)

// CircleArrowDown creates a Circle Arrow Down Lucide icon.
func CircleArrowDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-arrow-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M12 8v8"))),
		html.Child(html.SvgPath(html.AD("m8 12 4 4 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
