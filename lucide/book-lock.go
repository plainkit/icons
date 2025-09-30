package lucide

import (
	html "github.com/plainkit/html"
)

// BookLock creates a Book Lock Lucide icon.
func BookLock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-lock", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 6V4a2 2 0 1 0-4 0v2")),
		html.SvgPath(html.AD("M20 15v6a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20")),
		html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H10")),
		html.SvgRect(html.AWidth("8"), html.AHeight("5"), html.AX("12"), html.AY("6"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
