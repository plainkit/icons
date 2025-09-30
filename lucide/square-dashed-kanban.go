package lucide

import (
	html "github.com/plainkit/html"
)

// SquareDashedKanban creates a Square Dashed Kanban Lucide icon.
func SquareDashedKanban(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-dashed-kanban", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 7v7")),
		html.SvgPath(html.AD("M12 7v4")),
		html.SvgPath(html.AD("M16 7v9")),
		html.SvgPath(html.AD("M5 3a2 2 0 0 0-2 2")),
		html.SvgPath(html.AD("M9 3h1")),
		html.SvgPath(html.AD("M14 3h1")),
		html.SvgPath(html.AD("M19 3a2 2 0 0 1 2 2")),
		html.SvgPath(html.AD("M21 9v1")),
		html.SvgPath(html.AD("M21 14v1")),
		html.SvgPath(html.AD("M21 19a2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("M14 21h1")),
		html.SvgPath(html.AD("M9 21h1")),
		html.SvgPath(html.AD("M5 21a2 2 0 0 1-2-2")),
		html.SvgPath(html.AD("M3 14v1")),
		html.SvgPath(html.AD("M3 9v1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
