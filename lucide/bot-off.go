package lucide

import (
	html "github.com/plainkit/html"
)

// BotOff creates a Bot Off Lucide icon.
func BotOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bot-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13.67 8H18a2 2 0 0 1 2 2v4.33"))),
		html.Child(html.SvgPath(html.AD("M2 14h2"))),
		html.Child(html.SvgPath(html.AD("M20 14h2"))),
		html.Child(html.SvgPath(html.AD("M22 22 2 2"))),
		html.Child(html.SvgPath(html.AD("M8 8H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 1.414-.586"))),
		html.Child(html.SvgPath(html.AD("M9 13v2"))),
		html.Child(html.SvgPath(html.AD("M9.67 4H12v2.33"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
