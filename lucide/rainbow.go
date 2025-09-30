package lucide

import (
	html "github.com/plainkit/html"
)

// Rainbow creates a Rainbow Lucide icon.
func Rainbow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rainbow", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 17a10 10 0 0 0-20 0")),
		html.SvgPath(html.AD("M6 17a6 6 0 0 1 12 0")),
		html.SvgPath(html.AD("M10 17a2 2 0 0 1 4 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
