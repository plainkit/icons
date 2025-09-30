package lucide

import (
	html "github.com/plainkit/html"
)

// CircleQuestionMark creates a Circle Question Mark Lucide icon.
func CircleQuestionMark(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-question-mark", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3")),
		html.SvgPath(html.AD("M12 17h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
