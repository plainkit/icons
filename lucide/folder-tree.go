package lucide

import (
	html "github.com/plainkit/html"
)

// FolderTree creates a Folder Tree Lucide icon.
func FolderTree(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-tree", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 10a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-2.5a1 1 0 0 1-.8-.4l-.9-1.2A1 1 0 0 0 15 3h-2a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1Z"))),
		html.Child(html.SvgPath(html.AD("M20 21a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-2.9a1 1 0 0 1-.88-.55l-.42-.85a1 1 0 0 0-.92-.6H13a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1Z"))),
		html.Child(html.SvgPath(html.AD("M3 5a2 2 0 0 0 2 2h3"))),
		html.Child(html.SvgPath(html.AD("M3 3v13a2 2 0 0 0 2 2h3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
