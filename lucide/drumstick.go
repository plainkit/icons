package lucide

import (
	html "github.com/plainkit/html"
)

// Drumstick creates a Drumstick Lucide icon.
func Drumstick(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-drumstick", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15.4 15.63a7.875 6 135 1 1 6.23-6.23 4.5 3.43 135 0 0-6.23 6.23")),
		html.SvgPath(html.AD("m8.29 12.71-2.6 2.6a2.5 2.5 0 1 0-1.65 4.65A2.5 2.5 0 1 0 8.7 18.3l2.59-2.59")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
