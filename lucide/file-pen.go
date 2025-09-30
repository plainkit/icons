package lucide

import (
	html "github.com/plainkit/html"
)

// FilePen creates a File Pen Lucide icon.
func FilePen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-pen", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12.5 22H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v9.5")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M13.378 15.626a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
