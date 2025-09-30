package lucide

import (
	html "github.com/plainkit/html"
)

// Bell creates a Bell Lucide icon.
func Bell(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bell", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.268 21a2 2 0 0 0 3.464 0")),
		html.SvgPath(html.AD("M3.262 15.326A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673C19.41 13.956 18 12.499 18 8A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.738 7.326")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
