package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownAZ creates a Arrow Down A Z Lucide icon.
func ArrowDownAZ(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-a-z", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 16 4 4 4-4"))),
		html.Child(html.SvgPath(html.AD("M7 20V4"))),
		html.Child(html.SvgPath(html.AD("M20 8h-5"))),
		html.Child(html.SvgPath(html.AD("M15 10V6.5a2.5 2.5 0 0 1 5 0V10"))),
		html.Child(html.SvgPath(html.AD("M15 14h5l-5 6h5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
