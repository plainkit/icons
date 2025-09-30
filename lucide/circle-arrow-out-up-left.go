package lucide

import (
	html "github.com/plainkit/html"
)

// CircleArrowOutUpLeft creates a Circle Arrow Out Up Left Lucide icon.
func CircleArrowOutUpLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-arrow-out-up-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 8V2h6")),
		html.SvgPath(html.AD("m2 2 10 10")),
		html.SvgPath(html.AD("M12 2A10 10 0 1 1 2 12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
