package lucide

import (
	html "github.com/plainkit/html"
)

// FolderMinus creates a Folder Minus Lucide icon.
func FolderMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M9 13h6"))),
		html.Child(html.SvgPath(html.AD("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
