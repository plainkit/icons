package lucide

import (
	html "github.com/plainkit/html"
)

// Microwave creates a Microwave Lucide icon.
func Microwave(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-microwave", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("15"), html.AX("2"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("7"), html.AX("6"), html.AY("8"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("M18 8v7"))),
		html.Child(html.SvgPath(html.AD("M6 19v2"))),
		html.Child(html.SvgPath(html.AD("M18 19v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
