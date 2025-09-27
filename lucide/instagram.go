package lucide

import (
	html "github.com/plainkit/html"
)

// Instagram creates a Instagram Lucide icon.
func Instagram(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-instagram", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("20"), html.AX("2"), html.AY("2"), html.ARx("5"), html.ARy("5"))),
		html.Child(html.SvgPath(html.AD("M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"))),
		html.Child(html.SvgLine(html.AX1("17.5"), html.AX2("17.51"), html.AY1("6.5"), html.AY2("6.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
