package lucide

import (
	html "github.com/plainkit/html"
)

// FileType creates a File Type Lucide icon.
func FileType(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-type", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M9 13v-1h6v1")),
		html.SvgPath(html.AD("M12 12v6")),
		html.SvgPath(html.AD("M11 18h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
