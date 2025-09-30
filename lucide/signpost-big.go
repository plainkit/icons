package lucide

import (
	html "github.com/plainkit/html"
)

// SignpostBig creates a Signpost Big Lucide icon.
func SignpostBig(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-signpost-big", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 9H4L2 7l2-2h6")),
		html.SvgPath(html.AD("M14 5h6l2 2-2 2h-6")),
		html.SvgPath(html.AD("M10 22V4a2 2 0 1 1 4 0v18")),
		html.SvgPath(html.AD("M8 22h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
