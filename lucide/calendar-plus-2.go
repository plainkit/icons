package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarPlus2 creates a Calendar Plus 2 Lucide icon.
func CalendarPlus2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-plus-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 2v4")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M3 10h18")),
		html.SvgPath(html.AD("M10 16h4")),
		html.SvgPath(html.AD("M12 14v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
