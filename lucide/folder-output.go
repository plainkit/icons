package lucide

import (
	html "github.com/plainkit/html"
)

// FolderOutput creates a Folder Output Lucide icon.
func FolderOutput(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-output", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 7.5V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-1.5"))),
		html.Child(html.SvgPath(html.AD("M2 13h10"))),
		html.Child(html.SvgPath(html.AD("m5 10-3 3 3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
