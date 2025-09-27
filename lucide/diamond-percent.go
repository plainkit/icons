package lucide

import (
	html "github.com/plainkit/html"
)

// DiamondPercent creates a Diamond Percent Lucide icon.
func DiamondPercent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-diamond-percent", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0Z"))),
		html.Child(html.SvgPath(html.AD("M9.2 9.2h.01"))),
		html.Child(html.SvgPath(html.AD("m14.5 9.5-5 5"))),
		html.Child(html.SvgPath(html.AD("M14.7 14.8h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
