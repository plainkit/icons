package lucide

import (
	html "github.com/plainkit/html"
)

// StarOff creates a Star Off Lucide icon.
func StarOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-star-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8.34 8.34 2 9.27l5 4.87L5.82 21 12 17.77 18.18 21l-.59-3.43")),
		html.SvgPath(html.AD("M18.42 12.76 22 9.27l-6.91-1L12 2l-1.44 2.91")),
		html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
