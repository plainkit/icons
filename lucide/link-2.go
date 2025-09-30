package lucide

import (
	html "github.com/plainkit/html"
)

// Link2 creates a Link 2 Lucide icon.
func Link2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-link-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9 17H7A5 5 0 0 1 7 7h2")),
		html.SvgPath(html.AD("M15 7h2a5 5 0 1 1 0 10h-2")),
		html.SvgLine(html.AX1("8"), html.AX2("16"), html.AY1("12"), html.AY2("12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
