package lucide

import (
	html "github.com/plainkit/html"
)

// CircleMinus creates a Circle Minus Lucide icon.
func CircleMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-minus", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M8 12h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
