package lucide

import (
	html "github.com/plainkit/html"
)

// ListChecks creates a List Checks Lucide icon.
func ListChecks(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-checks", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 5h8")),
		html.SvgPath(html.AD("M13 12h8")),
		html.SvgPath(html.AD("M13 19h8")),
		html.SvgPath(html.AD("m3 17 2 2 4-4")),
		html.SvgPath(html.AD("m3 7 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
