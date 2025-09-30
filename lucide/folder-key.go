package lucide

import (
	html "github.com/plainkit/html"
)

// FolderKey creates a Folder Key Lucide icon.
func FolderKey(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-key", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("16"), html.ACy("20"), html.AR("2")),
		html.SvgPath(html.AD("M10 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v2")),
		html.SvgPath(html.AD("m22 14-4.5 4.5")),
		html.SvgPath(html.AD("m21 15 1 1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
