package lucide

import (
	html "github.com/plainkit/html"
)

// CandyOff creates a Candy Off Lucide icon.
func CandyOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-candy-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 10v7.9"))),
		html.Child(html.SvgPath(html.AD("M11.802 6.145a5 5 0 0 1 6.053 6.053"))),
		html.Child(html.SvgPath(html.AD("M14 6.1v2.243"))),
		html.Child(html.SvgPath(html.AD("m15.5 15.571-.964.964a5 5 0 0 1-7.071 0 5 5 0 0 1 0-7.07l.964-.965"))),
		html.Child(html.SvgPath(html.AD("M16 7V3a1 1 0 0 1 1.707-.707 2.5 2.5 0 0 0 2.152.717 1 1 0 0 1 1.131 1.131 2.5 2.5 0 0 0 .717 2.152A1 1 0 0 1 21 8h-4"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M8 17v4a1 1 0 0 1-1.707.707 2.5 2.5 0 0 0-2.152-.717 1 1 0 0 1-1.131-1.131 2.5 2.5 0 0 0-.717-2.152A1 1 0 0 1 3 16h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
