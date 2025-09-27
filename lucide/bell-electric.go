package lucide

import (
	html "github.com/plainkit/html"
)

// BellElectric creates a Bell Electric Lucide icon.
func BellElectric(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bell-electric", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18.518 17.347A7 7 0 0 1 14 19"))),
		html.Child(html.SvgPath(html.AD("M18.8 4A11 11 0 0 1 20 9"))),
		html.Child(html.SvgPath(html.AD("M9 9h.01"))),
		html.Child(html.SvgCircle(html.ACx("20"), html.ACy("16"), html.AR("2"))),
		html.Child(html.SvgCircle(html.ACx("9"), html.ACy("9"), html.AR("7"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("6"), html.AX("4"), html.AY("16"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
