package lucide

import (
	html "github.com/plainkit/html"
)

// BadgePoundSterling creates a Badge Pound Sterling Lucide icon.
func BadgePoundSterling(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-pound-sterling", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z")),
		html.SvgPath(html.AD("M8 12h4")),
		html.SvgPath(html.AD("M10 16V9.5a2.5 2.5 0 0 1 5 0")),
		html.SvgPath(html.AD("M8 16h7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
