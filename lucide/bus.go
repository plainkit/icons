package lucide

import (
	html "github.com/plainkit/html"
)

// Bus creates a Bus Lucide icon.
func Bus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 6v6"))),
		html.Child(html.SvgPath(html.AD("M15 6v6"))),
		html.Child(html.SvgPath(html.AD("M2 12h19.6"))),
		html.Child(html.SvgPath(html.AD("M18 18h3s.5-1.7.8-2.8c.1-.4.2-.8.2-1.2 0-.4-.1-.8-.2-1.2l-1.4-5C20.1 6.8 19.1 6 18 6H4a2 2 0 0 0-2 2v10h3"))),
		html.Child(html.SvgCircle(html.ACx("7"), html.ACy("18"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M9 18h5"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("18"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
