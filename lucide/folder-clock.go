package lucide

import (
	html "github.com/plainkit/html"
)

// FolderClock creates a Folder Clock Lucide icon.
func FolderClock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-clock", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 14v2.2l1.6 1"))),
		html.Child(html.SvgPath(html.AD("M7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("16"), html.AR("6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
