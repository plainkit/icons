package lucide

import (
	html "github.com/plainkit/html"
)

// CableCar creates a Cable Car Lucide icon.
func CableCar(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cable-car", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 3h.01"))),
		html.Child(html.SvgPath(html.AD("M14 2h.01"))),
		html.Child(html.SvgPath(html.AD("m2 9 20-5"))),
		html.Child(html.SvgPath(html.AD("M12 12V6.5"))),
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("10"), html.AX("4"), html.AY("12"), html.ARx("3"))),
		html.Child(html.SvgPath(html.AD("M9 12v5"))),
		html.Child(html.SvgPath(html.AD("M15 12v5"))),
		html.Child(html.SvgPath(html.AD("M4 17h16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
