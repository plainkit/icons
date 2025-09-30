package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowBigDownDash creates a Arrow Big Down Dash Lucide icon.
func ArrowBigDownDash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-big-down-dash", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 11a1 1 0 0 0 1 1h2.939a1 1 0 0 1 .75 1.811l-6.835 6.836a1.207 1.207 0 0 1-1.707 0L4.31 13.81a1 1 0 0 1 .75-1.811H8a1 1 0 0 0 1-1V9a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1z")),
		html.SvgPath(html.AD("M9 4h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
