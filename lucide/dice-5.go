package lucide

import (
	html "github.com/plainkit/html"
)

// Dice5 creates a Dice 5 Lucide icon.
func Dice5(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dice-5", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("M16 8h.01"))),
		html.Child(html.SvgPath(html.AD("M8 8h.01"))),
		html.Child(html.SvgPath(html.AD("M8 16h.01"))),
		html.Child(html.SvgPath(html.AD("M16 16h.01"))),
		html.Child(html.SvgPath(html.AD("M12 12h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
