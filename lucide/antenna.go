package lucide

import (
	html "github.com/plainkit/html"
)

// Antenna creates a Antenna Lucide icon.
func Antenna(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-antenna", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 12 7 2")),
		html.SvgPath(html.AD("m7 12 5-10")),
		html.SvgPath(html.AD("m12 12 5-10")),
		html.SvgPath(html.AD("m17 12 5-10")),
		html.SvgPath(html.AD("M4.5 7h15")),
		html.SvgPath(html.AD("M12 16v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
