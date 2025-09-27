package lucide

import (
	html "github.com/plainkit/html"
)

// FolderSync creates a Folder Sync Lucide icon.
func FolderSync(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-sync", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v.5"))),
		html.Child(html.SvgPath(html.AD("M12 10v4h4"))),
		html.Child(html.SvgPath(html.AD("m12 14 1.535-1.605a5 5 0 0 1 8 1.5"))),
		html.Child(html.SvgPath(html.AD("M22 22v-4h-4"))),
		html.Child(html.SvgPath(html.AD("m22 18-1.535 1.605a5 5 0 0 1-8-1.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
