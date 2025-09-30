package lucide

import (
	html "github.com/plainkit/html"
)

// SquareKanban creates a Square Kanban Lucide icon.
func SquareKanban(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-kanban", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M8 7v7")),
		html.SvgPath(html.AD("M12 7v4")),
		html.SvgPath(html.AD("M16 7v9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
