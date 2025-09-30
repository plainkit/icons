package lucide

import (
	html "github.com/plainkit/html"
)

// FileDiff creates a File Diff Lucide icon.
func FileDiff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-diff", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgPath(html.AD("M9 10h6")),
		html.SvgPath(html.AD("M12 13V7")),
		html.SvgPath(html.AD("M9 17h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
