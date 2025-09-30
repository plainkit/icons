package lucide

import (
	html "github.com/plainkit/html"
)

// Ghost creates a Ghost Lucide icon.
func Ghost(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ghost", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9 10h.01")),
		html.SvgPath(html.AD("M15 10h.01")),
		html.SvgPath(html.AD("M12 2a8 8 0 0 0-8 8v12l3-3 2.5 2.5L12 19l2.5 2.5L17 19l3 3V10a8 8 0 0 0-8-8z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
