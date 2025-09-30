package lucide

import (
	html "github.com/plainkit/html"
)

// GitBranchPlus creates a Git Branch Plus Lucide icon.
func GitBranchPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-branch-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 3v12")),
		html.SvgPath(html.AD("M18 9a3 3 0 1 0 0-6 3 3 0 0 0 0 6z")),
		html.SvgPath(html.AD("M6 21a3 3 0 1 0 0-6 3 3 0 0 0 0 6z")),
		html.SvgPath(html.AD("M15 6a9 9 0 0 0-9 9")),
		html.SvgPath(html.AD("M18 15v6")),
		html.SvgPath(html.AD("M21 18h-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
