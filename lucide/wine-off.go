package lucide

import (
	html "github.com/plainkit/html"
)

// WineOff creates a Wine Off Lucide icon.
func WineOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wine-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 22h8")),
		html.SvgPath(html.AD("M7 10h3m7 0h-1.343")),
		html.SvgPath(html.AD("M12 15v7")),
		html.SvgPath(html.AD("M7.307 7.307A12.33 12.33 0 0 0 7 10a5 5 0 0 0 7.391 4.391M8.638 2.981C8.75 2.668 8.872 2.34 9 2h6c1.5 4 2 6 2 8 0 .407-.05.809-.145 1.198")),
		html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
