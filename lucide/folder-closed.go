package lucide

import (
	html "github.com/plainkit/html"
)

// FolderClosed creates a Folder Closed Lucide icon.
func FolderClosed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-closed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"))),
		html.Child(html.SvgPath(html.AD("M2 10h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
