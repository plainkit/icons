package lucide

import (
	html "github.com/plainkit/html"
)

// Trophy creates a Trophy Lucide icon.
func Trophy(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-trophy", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 14.66v1.626a2 2 0 0 1-.976 1.696A5 5 0 0 0 7 21.978")),
		html.SvgPath(html.AD("M14 14.66v1.626a2 2 0 0 0 .976 1.696A5 5 0 0 1 17 21.978")),
		html.SvgPath(html.AD("M18 9h1.5a1 1 0 0 0 0-5H18")),
		html.SvgPath(html.AD("M4 22h16")),
		html.SvgPath(html.AD("M6 9a6 6 0 0 0 12 0V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1z")),
		html.SvgPath(html.AD("M6 9H4.5a1 1 0 0 1 0-5H6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
