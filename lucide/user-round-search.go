package lucide

import (
	html "github.com/plainkit/html"
)

// UserRoundSearch creates a User Round Search Lucide icon.
func UserRoundSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-round-search", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("10"), html.ACy("8"), html.AR("5")),
		html.SvgPath(html.AD("M2 21a8 8 0 0 1 10.434-7.62")),
		html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3")),
		html.SvgPath(html.AD("m22 22-1.9-1.9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
