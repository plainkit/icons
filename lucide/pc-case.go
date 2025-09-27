package lucide

import (
	html "github.com/plainkit/html"
)

// PcCase creates a Pc Case Lucide icon.
func PcCase(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pc-case", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("20"), html.AX("5"), html.AY("2"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M15 14h.01"))),
		html.Child(html.SvgPath(html.AD("M9 6h6"))),
		html.Child(html.SvgPath(html.AD("M9 10h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
