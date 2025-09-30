package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowBigLeft creates a Arrow Big Left Lucide icon.
func ArrowBigLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-big-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 9a1 1 0 0 1-1-1V5.061a1 1 0 0 0-1.811-.75l-6.835 6.836a1.207 1.207 0 0 0 0 1.707l6.835 6.835a1 1 0 0 0 1.811-.75V16a1 1 0 0 1 1-1h6a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
