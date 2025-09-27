package lucide

import (
	html "github.com/plainkit/html"
)

// FolderOpen creates a Folder Open Lucide icon.
func FolderOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-open", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m6 14 1.5-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.54 6a2 2 0 0 1-1.95 1.5H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H18a2 2 0 0 1 2 2v2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
