package lucide

import (
	html "github.com/plainkit/html"
)

// ListVideo creates a List Video Lucide icon.
func ListVideo(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-video", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 5H3")),
		html.SvgPath(html.AD("M10 12H3")),
		html.SvgPath(html.AD("M10 19H3")),
		html.SvgPath(html.AD("M15 12.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
