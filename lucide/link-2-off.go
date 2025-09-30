package lucide

import (
	html "github.com/plainkit/html"
)

// Link2Off creates a Link 2 Off Lucide icon.
func Link2Off(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-link-2-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9 17H7A5 5 0 0 1 7 7")),
		html.SvgPath(html.AD("M15 7h2a5 5 0 0 1 4 8")),
		html.SvgLine(html.AX1("8"), html.AX2("12"), html.AY1("12"), html.AY2("12")),
		html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
