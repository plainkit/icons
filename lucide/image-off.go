package lucide

import (
	html "github.com/plainkit/html"
)

// ImageOff creates a Image Off Lucide icon.
func ImageOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-image-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22"))),
		html.Child(html.SvgPath(html.AD("M10.41 10.41a2 2 0 1 1-2.83-2.83"))),
		html.Child(html.SvgLine(html.AX1("13.5"), html.AX2("6"), html.AY1("13.5"), html.AY2("21"))),
		html.Child(html.SvgLine(html.AX1("18"), html.AX2("21"), html.AY1("12"), html.AY2("15"))),
		html.Child(html.SvgPath(html.AD("M3.59 3.59A1.99 1.99 0 0 0 3 5v14a2 2 0 0 0 2 2h14c.55 0 1.052-.22 1.41-.59"))),
		html.Child(html.SvgPath(html.AD("M21 15V5a2 2 0 0 0-2-2H9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
