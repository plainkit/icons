package lucide

import (
	html "github.com/plainkit/html"
)

// CirclePoundSterling creates a Circle Pound Sterling Lucide icon.
func CirclePoundSterling(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-pound-sterling", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 16V9.5a1 1 0 0 1 5 0")),
		html.SvgPath(html.AD("M8 12h4")),
		html.SvgPath(html.AD("M8 16h7")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
