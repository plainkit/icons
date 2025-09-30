package lucide

import (
	html "github.com/plainkit/html"
)

// FileSymlink creates a File Symlink Lucide icon.
func FileSymlink(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-symlink", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m10 18 3-3-3-3")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M4 11V4a2 2 0 0 1 2-2h9l5 5v13a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
