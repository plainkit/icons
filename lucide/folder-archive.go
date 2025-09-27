package lucide

import (
	html "github.com/plainkit/html"
)

// FolderArchive creates a Folder Archive Lucide icon.
func FolderArchive(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-archive", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("15"), html.ACy("19"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M20.9 19.8A2 2 0 0 0 22 18V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h5.1"))),
		html.Child(html.SvgPath(html.AD("M15 11v-1"))),
		html.Child(html.SvgPath(html.AD("M15 17v-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
