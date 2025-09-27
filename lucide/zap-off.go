package lucide

import (
	html "github.com/plainkit/html"
)

// ZapOff creates a Zap Off Lucide icon.
func ZapOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-zap-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.513 4.856 13.12 2.17a.5.5 0 0 1 .86.46l-1.377 4.317"))),
		html.Child(html.SvgPath(html.AD("M15.656 10H20a1 1 0 0 1 .78 1.63l-1.72 1.773"))),
		html.Child(html.SvgPath(html.AD("M16.273 16.273 10.88 21.83a.5.5 0 0 1-.86-.46l1.92-6.02A1 1 0 0 0 11 14H4a1 1 0 0 1-.78-1.63l4.507-4.643"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
