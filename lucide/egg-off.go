package lucide

import (
	html "github.com/plainkit/html"
)

// EggOff creates a Egg Off Lucide icon.
func EggOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-egg-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M20 14.347V14c0-6-4-12-8-12-1.078 0-2.157.436-3.157 1.19"))),
		html.Child(html.SvgPath(html.AD("M6.206 6.21C4.871 8.4 4 11.2 4 14a8 8 0 0 0 14.568 4.568"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
