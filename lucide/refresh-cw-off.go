package lucide

import (
	html "github.com/plainkit/html"
)

// RefreshCwOff creates a Refresh Cw Off Lucide icon.
func RefreshCwOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-refresh-cw-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 8L18.74 5.74A9.75 9.75 0 0 0 12 3C11 3 10.03 3.16 9.13 3.47"))),
		html.Child(html.SvgPath(html.AD("M8 16H3v5"))),
		html.Child(html.SvgPath(html.AD("M3 12C3 9.51 4 7.26 5.64 5.64"))),
		html.Child(html.SvgPath(html.AD("m3 16 2.26 2.26A9.75 9.75 0 0 0 12 21c2.49 0 4.74-1 6.36-2.64"))),
		html.Child(html.SvgPath(html.AD("M21 12c0 1-.16 1.97-.47 2.87"))),
		html.Child(html.SvgPath(html.AD("M21 3v5h-5"))),
		html.Child(html.SvgPath(html.AD("M22 22 2 2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
