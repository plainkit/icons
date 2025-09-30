package lucide

import (
	html "github.com/plainkit/html"
)

// FolderRoot creates a Folder Root Lucide icon.
func FolderRoot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-root", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z")),
		html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("2")),
		html.SvgPath(html.AD("M12 15v5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
