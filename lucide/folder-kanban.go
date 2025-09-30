package lucide

import (
	html "github.com/plainkit/html"
)

// FolderKanban creates a Folder Kanban Lucide icon.
func FolderKanban(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-kanban", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z")),
		html.SvgPath(html.AD("M8 10v4")),
		html.SvgPath(html.AD("M12 10v2")),
		html.SvgPath(html.AD("M16 10v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
