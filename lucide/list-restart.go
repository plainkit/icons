package lucide

import (
	html "github.com/plainkit/html"
)

// ListRestart creates a List Restart Lucide icon.
func ListRestart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-restart", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 5H3")),
		html.SvgPath(html.AD("M7 12H3")),
		html.SvgPath(html.AD("M7 19H3")),
		html.SvgPath(html.AD("M12 18a5 5 0 0 0 9-3 4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L11 14")),
		html.SvgPath(html.AD("M11 10v4h4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
