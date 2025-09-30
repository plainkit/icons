package lucide

import (
	html "github.com/plainkit/html"
)

// AlignCenterVertical creates a Align Center Vertical Lucide icon.
func AlignCenterVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-center-vertical", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 2v20")),
		html.SvgPath(html.AD("M8 10H4a2 2 0 0 1-2-2V6c0-1.1.9-2 2-2h4")),
		html.SvgPath(html.AD("M16 10h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4")),
		html.SvgPath(html.AD("M8 20H7a2 2 0 0 1-2-2v-2c0-1.1.9-2 2-2h1")),
		html.SvgPath(html.AD("M16 14h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
