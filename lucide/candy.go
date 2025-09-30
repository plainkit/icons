package lucide

import (
	html "github.com/plainkit/html"
)

// Candy creates a Candy Lucide icon.
func Candy(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-candy", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 7v10.9")),
		html.SvgPath(html.AD("M14 6.1V17")),
		html.SvgPath(html.AD("M16 7V3a1 1 0 0 1 1.707-.707 2.5 2.5 0 0 0 2.152.717 1 1 0 0 1 1.131 1.131 2.5 2.5 0 0 0 .717 2.152A1 1 0 0 1 21 8h-4")),
		html.SvgPath(html.AD("M16.536 7.465a5 5 0 0 0-7.072 0l-2 2a5 5 0 0 0 0 7.07 5 5 0 0 0 7.072 0l2-2a5 5 0 0 0 0-7.07")),
		html.SvgPath(html.AD("M8 17v4a1 1 0 0 1-1.707.707 2.5 2.5 0 0 0-2.152-.717 1 1 0 0 1-1.131-1.131 2.5 2.5 0 0 0-.717-2.152A1 1 0 0 1 3 16h4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
