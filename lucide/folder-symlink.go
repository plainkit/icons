package lucide

import (
	html "github.com/plainkit/html"
)

// FolderSymlink creates a Folder Symlink Lucide icon.
func FolderSymlink(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-symlink", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 9.35V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h7"))),
		html.Child(html.SvgPath(html.AD("m8 16 3-3-3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
