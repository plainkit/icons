package lucide

import (
	html "github.com/plainkit/html"
)

// CircleCheck creates a Circle Check Lucide icon.
func CircleCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("m9 12 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
