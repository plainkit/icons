package lucide

import (
	html "github.com/plainkit/html"
)

// CircleEllipsis creates a Circle Ellipsis Lucide icon.
func CircleEllipsis(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-ellipsis", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M17 12h.01")),
		html.SvgPath(html.AD("M12 12h.01")),
		html.SvgPath(html.AD("M7 12h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
