package lucide

import (
	html "github.com/plainkit/html"
)

// ThumbsDown creates a Thumbs Down Lucide icon.
func ThumbsDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-thumbs-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 14V2")),
		html.SvgPath(html.AD("M9 18.12 10 14H4.17a2 2 0 0 1-1.92-2.56l2.33-8A2 2 0 0 1 6.5 2H20a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2.76a2 2 0 0 0-1.79 1.11L12 22a3.13 3.13 0 0 1-3-3.88Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
