package lucide

import (
	html "github.com/plainkit/html"
)

// BotMessageSquare creates a Bot Message Square Lucide icon.
func BotMessageSquare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bot-message-square", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 6V2H8"))),
		html.Child(html.SvgPath(html.AD("M15 11v2"))),
		html.Child(html.SvgPath(html.AD("M2 12h2"))),
		html.Child(html.SvgPath(html.AD("M20 12h2"))),
		html.Child(html.SvgPath(html.AD("M20 16a2 2 0 0 1-2 2H8.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 4 20.286V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2z"))),
		html.Child(html.SvgPath(html.AD("M9 11v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
