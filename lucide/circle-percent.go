package lucide

import (
	html "github.com/plainkit/html"
)

// CirclePercent creates a Circle Percent Lucide icon.
func CirclePercent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-percent", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("m15 9-6 6"))),
		html.Child(html.SvgPath(html.AD("M9 9h.01"))),
		html.Child(html.SvgPath(html.AD("M15 15h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
