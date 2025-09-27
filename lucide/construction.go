package lucide

import (
	html "github.com/plainkit/html"
)

// Construction creates a Construction Lucide icon.
func Construction(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-construction", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("8"), html.AX("2"), html.AY("6"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("M17 14v7"))),
		html.Child(html.SvgPath(html.AD("M7 14v7"))),
		html.Child(html.SvgPath(html.AD("M17 3v3"))),
		html.Child(html.SvgPath(html.AD("M7 3v3"))),
		html.Child(html.SvgPath(html.AD("M10 14 2.3 6.3"))),
		html.Child(html.SvgPath(html.AD("m14 6 7.7 7.7"))),
		html.Child(html.SvgPath(html.AD("m8 6 8 8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
