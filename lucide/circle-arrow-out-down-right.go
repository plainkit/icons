package lucide

import (
	html "github.com/plainkit/html"
)

// CircleArrowOutDownRight creates a Circle Arrow Out Down Right Lucide icon.
func CircleArrowOutDownRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-arrow-out-down-right", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 22a10 10 0 1 1 10-10")),
		html.SvgPath(html.AD("M22 22 12 12")),
		html.SvgPath(html.AD("M22 16v6h-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
