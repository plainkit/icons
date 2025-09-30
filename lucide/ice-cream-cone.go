package lucide

import (
	html "github.com/plainkit/html"
)

// IceCreamCone creates a Ice Cream Cone Lucide icon.
func IceCreamCone(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ice-cream-cone", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m7 11 4.08 10.35a1 1 0 0 0 1.84 0L17 11")),
		html.SvgPath(html.AD("M17 7A5 5 0 0 0 7 7")),
		html.SvgPath(html.AD("M17 7a2 2 0 0 1 0 4H7a2 2 0 0 1 0-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
