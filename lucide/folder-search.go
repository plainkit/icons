package lucide

import (
	html "github.com/plainkit/html"
)

// FolderSearch creates a Folder Search Lucide icon.
func FolderSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-search", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v4.1")),
		html.SvgPath(html.AD("m21 21-1.9-1.9")),
		html.SvgCircle(html.ACx("17"), html.ACy("17"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
