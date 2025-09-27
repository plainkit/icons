package lucide

import (
	html "github.com/plainkit/html"
)

// RotateCw creates a Rotate Cw Lucide icon.
func RotateCw(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rotate-cw", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 12a9 9 0 1 1-9-9c2.52 0 4.93 1 6.74 2.74L21 8"))),
		html.Child(html.SvgPath(html.AD("M21 3v5h-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
