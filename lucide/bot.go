package lucide

import (
	html "github.com/plainkit/html"
)

// Bot creates a Bot Lucide icon.
func Bot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bot", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 8V4H8"))),
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("12"), html.AX("4"), html.AY("8"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M2 14h2"))),
		html.Child(html.SvgPath(html.AD("M20 14h2"))),
		html.Child(html.SvgPath(html.AD("M15 13v2"))),
		html.Child(html.SvgPath(html.AD("M9 13v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
