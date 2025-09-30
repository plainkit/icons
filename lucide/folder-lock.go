package lucide

import (
	html "github.com/plainkit/html"
)

// FolderLock creates a Folder Lock Lucide icon.
func FolderLock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-lock", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("5"), html.AX("14"), html.AY("17"), html.ARx("1")),
		html.SvgPath(html.AD("M10 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v2.5")),
		html.SvgPath(html.AD("M20 17v-2a2 2 0 1 0-4 0v2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
