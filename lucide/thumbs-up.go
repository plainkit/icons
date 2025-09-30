package lucide

import (
	html "github.com/plainkit/html"
)

// ThumbsUp creates a Thumbs Up Lucide icon.
func ThumbsUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-thumbs-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M7 10v12")),
		html.SvgPath(html.AD("M15 5.88 14 10h5.83a2 2 0 0 1 1.92 2.56l-2.33 8A2 2 0 0 1 17.5 22H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h2.76a2 2 0 0 0 1.79-1.11L12 2a3.13 3.13 0 0 1 3 3.88Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
