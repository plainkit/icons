package lucide

import (
	html "github.com/plainkit/html"
)

// MountainSnow creates a Mountain Snow Lucide icon.
func MountainSnow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mountain-snow", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m8 3 4 8 5-5 5 15H2L8 3z")),
		html.SvgPath(html.AD("M4.14 15.08c2.62-1.57 5.24-1.43 7.86.42 2.74 1.94 5.49 2 8.23.19")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
