package lucide

import (
	html "github.com/plainkit/html"
)

// FolderGit2 creates a Folder Git 2 Lucide icon.
func FolderGit2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-git-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v5")),
		html.SvgCircle(html.ACx("13"), html.ACy("12"), html.AR("2")),
		html.SvgPath(html.AD("M18 19c-2.8 0-5-2.2-5-5v8")),
		html.SvgCircle(html.ACx("20"), html.ACy("19"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
