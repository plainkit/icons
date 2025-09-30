package lucide

import (
	html "github.com/plainkit/html"
)

// Cross creates a Cross Lucide icon.
func Cross(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cross", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 9a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h4a1 1 0 0 1 1 1v4a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-4a1 1 0 0 1 1-1h4a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-4a1 1 0 0 1-1-1V4a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4a1 1 0 0 1-1 1z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
