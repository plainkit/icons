package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowBigUpDash creates a Arrow Big Up Dash Lucide icon.
func ArrowBigUpDash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-big-up-dash", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9 13a1 1 0 0 0-1-1H5.061a1 1 0 0 1-.75-1.811l6.836-6.835a1.207 1.207 0 0 1 1.707 0l6.835 6.835a1 1 0 0 1-.75 1.811H16a1 1 0 0 0-1 1v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1z")),
		html.SvgPath(html.AD("M9 20h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
