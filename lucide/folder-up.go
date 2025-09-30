package lucide

import (
	html "github.com/plainkit/html"
)

// FolderUp creates a Folder Up Lucide icon.
func FolderUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z")),
		html.SvgPath(html.AD("M12 10v6")),
		html.SvgPath(html.AD("m9 13 3-3 3 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
