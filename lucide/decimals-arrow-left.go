package lucide

import (
	html "github.com/plainkit/html"
)

// DecimalsArrowLeft creates a Decimals Arrow Left Lucide icon.
func DecimalsArrowLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-decimals-arrow-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m13 21-3-3 3-3"))),
		html.Child(html.SvgPath(html.AD("M20 18H10"))),
		html.Child(html.SvgPath(html.AD("M3 11h.01"))),
		html.Child(html.SvgRect(html.AWidth("5"), html.AHeight("8"), html.AX("6"), html.AY("3"), html.ARx("2.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
