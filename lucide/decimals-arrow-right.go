package lucide

import (
	html "github.com/plainkit/html"
)

// DecimalsArrowRight creates a Decimals Arrow Right Lucide icon.
func DecimalsArrowRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-decimals-arrow-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 18h10"))),
		html.Child(html.SvgPath(html.AD("m17 21 3-3-3-3"))),
		html.Child(html.SvgPath(html.AD("M3 11h.01"))),
		html.Child(html.SvgRect(html.AWidth("5"), html.AHeight("8"), html.AX("15"), html.AY("3"), html.ARx("2.5"))),
		html.Child(html.SvgRect(html.AWidth("5"), html.AHeight("8"), html.AX("6"), html.AY("3"), html.ARx("2.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
