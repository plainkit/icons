package lucide

import (
	html "github.com/plainkit/html"
)

// GitCompare creates a Git Compare Lucide icon.
func GitCompare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-compare", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("6"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M13 6h3a2 2 0 0 1 2 2v7"))),
		html.Child(html.SvgPath(html.AD("M11 18H8a2 2 0 0 1-2-2V9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
