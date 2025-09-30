package lucide

import (
	html "github.com/plainkit/html"
)

// PocketKnife creates a Pocket Knife Lucide icon.
func PocketKnife(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pocket-knife", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 2v1c0 1 2 1 2 2S3 6 3 7s2 1 2 2-2 1-2 2 2 1 2 2")),
		html.SvgPath(html.AD("M18 6h.01")),
		html.SvgPath(html.AD("M6 18h.01")),
		html.SvgPath(html.AD("M20.83 8.83a4 4 0 0 0-5.66-5.66l-12 12a4 4 0 1 0 5.66 5.66Z")),
		html.SvgPath(html.AD("M18 11.66V22a4 4 0 0 0 4-4V6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
