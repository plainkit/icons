package lucide

import (
	html "github.com/plainkit/html"
)

// Kanban creates a Kanban Lucide icon.
func Kanban(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-kanban", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 3v14")),
		html.SvgPath(html.AD("M12 3v8")),
		html.SvgPath(html.AD("M19 3v18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
