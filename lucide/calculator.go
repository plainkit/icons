package lucide

import (
	html "github.com/plainkit/html"
)

// Calculator creates a Calculator Lucide icon.
func Calculator(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calculator", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("16"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2")),
		html.SvgLine(html.AX1("8"), html.AX2("16"), html.AY1("6"), html.AY2("6")),
		html.SvgLine(html.AX1("16"), html.AX2("16"), html.AY1("14"), html.AY2("18")),
		html.SvgPath(html.AD("M16 10h.01")),
		html.SvgPath(html.AD("M12 10h.01")),
		html.SvgPath(html.AD("M8 10h.01")),
		html.SvgPath(html.AD("M12 14h.01")),
		html.SvgPath(html.AD("M8 14h.01")),
		html.SvgPath(html.AD("M12 18h.01")),
		html.SvgPath(html.AD("M8 18h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
