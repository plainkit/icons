package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeQuestionMark creates a Badge Question Mark Lucide icon.
func BadgeQuestionMark(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-question-mark", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z")),
		html.SvgPath(html.AD("M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3")),
		html.SvgLine(html.AX1("12"), html.AX2("12.01"), html.AY1("17"), html.AY2("17")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
