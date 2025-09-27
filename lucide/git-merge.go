package lucide

import (
	html "github.com/plainkit/html"
)

// GitMerge creates a Git Merge Lucide icon.
func GitMerge(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-merge", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("6"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M6 21V9a9 9 0 0 0 9 9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
