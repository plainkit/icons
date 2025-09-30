package lucide

import (
	html "github.com/plainkit/html"
)

// IceCreamBowl creates a Ice Cream Bowl Lucide icon.
func IceCreamBowl(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ice-cream-bowl", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 17c5 0 8-2.69 8-6H4c0 3.31 3 6 8 6m-4 4h8m-4-3v3M5.14 11a3.5 3.5 0 1 1 6.71 0")),
		html.SvgPath(html.AD("M12.14 11a3.5 3.5 0 1 1 6.71 0")),
		html.SvgPath(html.AD("M15.5 6.5a3.5 3.5 0 1 0-7 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
