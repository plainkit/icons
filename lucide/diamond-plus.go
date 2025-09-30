package lucide

import (
	html "github.com/plainkit/html"
)

// DiamondPlus creates a Diamond Plus Lucide icon.
func DiamondPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-diamond-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 8v8")),
		html.SvgPath(html.AD("M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0z")),
		html.SvgPath(html.AD("M8 12h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
