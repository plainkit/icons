package lucide

import (
	html "github.com/plainkit/html"
)

// BellPlus creates a Bell Plus Lucide icon.
func BellPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bell-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.268 21a2 2 0 0 0 3.464 0")),
		html.SvgPath(html.AD("M15 8h6")),
		html.SvgPath(html.AD("M18 5v6")),
		html.SvgPath(html.AD("M20.002 14.464a9 9 0 0 0 .738.863A1 1 0 0 1 20 17H4a1 1 0 0 1-.74-1.673C4.59 13.956 6 12.499 6 8a6 6 0 0 1 8.75-5.332")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
