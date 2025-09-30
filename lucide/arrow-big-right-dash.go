package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowBigRightDash creates a Arrow Big Right Dash Lucide icon.
func ArrowBigRightDash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-big-right-dash", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 9a1 1 0 0 0 1-1V5.061a1 1 0 0 1 1.811-.75l6.836 6.836a1.207 1.207 0 0 1 0 1.707l-6.836 6.835a1 1 0 0 1-1.811-.75V16a1 1 0 0 0-1-1H9a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z")),
		html.SvgPath(html.AD("M4 9v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
