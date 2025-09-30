package lucide

import (
	html "github.com/plainkit/html"
)

// BookAlert creates a Book Alert Lucide icon.
func BookAlert(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-alert", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 13h.01")),
		html.SvgPath(html.AD("M12 6v3")),
		html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
