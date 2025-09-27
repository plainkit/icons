package lucide

import (
	html "github.com/plainkit/html"
)

// Dice2 creates a Dice 2 Lucide icon.
func Dice2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dice-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("M15 9h.01"))),
		html.Child(html.SvgPath(html.AD("M9 15h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
