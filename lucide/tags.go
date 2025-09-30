package lucide

import (
	html "github.com/plainkit/html"
)

// Tags creates a Tags Lucide icon.
func Tags(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tags", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13.172 2a2 2 0 0 1 1.414.586l6.71 6.71a2.4 2.4 0 0 1 0 3.408l-4.592 4.592a2.4 2.4 0 0 1-3.408 0l-6.71-6.71A2 2 0 0 1 6 9.172V3a1 1 0 0 1 1-1z")),
		html.SvgPath(html.AD("M2 7v6.172a2 2 0 0 0 .586 1.414l6.71 6.71a2.4 2.4 0 0 0 3.191.193")),
		html.SvgCircle(html.ACx("10.5"), html.ACy("6.5"), html.AR(".5"), html.AFill("currentColor")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
