package lucide

import (
	html "github.com/plainkit/html"
)

// FolderGit creates a Folder Git Lucide icon.
func FolderGit(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-git", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"))),
		html.Child(html.SvgPath(html.AD("M14 13h3"))),
		html.Child(html.SvgPath(html.AD("M7 13h3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
