package lucide

import (
	html "github.com/plainkit/html"
)

// IdCard creates a Id Card Lucide icon.
func IdCard(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-id-card", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 10h2")),
		html.SvgPath(html.AD("M16 14h2")),
		html.SvgPath(html.AD("M6.17 15a3 3 0 0 1 5.66 0")),
		html.SvgCircle(html.ACx("9"), html.ACy("11"), html.AR("2")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("5"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
