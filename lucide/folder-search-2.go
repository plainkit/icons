package lucide

import (
	html "github.com/plainkit/html"
)

// FolderSearch2 creates a Folder Search 2 Lucide icon.
func FolderSearch2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-search-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("11.5"), html.ACy("12.5"), html.AR("2.5"))),
		html.Child(html.SvgPath(html.AD("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"))),
		html.Child(html.SvgPath(html.AD("M13.3 14.3 15 16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
