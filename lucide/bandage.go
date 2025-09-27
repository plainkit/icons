package lucide

import (
	html "github.com/plainkit/html"
)

// Bandage creates a Bandage Lucide icon.
func Bandage(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bandage", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 10.01h.01"))),
		html.Child(html.SvgPath(html.AD("M10 14.01h.01"))),
		html.Child(html.SvgPath(html.AD("M14 10.01h.01"))),
		html.Child(html.SvgPath(html.AD("M14 14.01h.01"))),
		html.Child(html.SvgPath(html.AD("M18 6v11.5"))),
		html.Child(html.SvgPath(html.AD("M6 6v12"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
