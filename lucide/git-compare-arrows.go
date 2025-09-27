package lucide

import (
	html "github.com/plainkit/html"
)

// GitCompareArrows creates a Git Compare Arrows Lucide icon.
func GitCompareArrows(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-compare-arrows", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("5"), html.ACy("6"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M12 6h5a2 2 0 0 1 2 2v7"))),
		html.Child(html.SvgPath(html.AD("m15 9-3-3 3-3"))),
		html.Child(html.SvgCircle(html.ACx("19"), html.ACy("18"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M12 18H7a2 2 0 0 1-2-2V9"))),
		html.Child(html.SvgPath(html.AD("m9 15 3 3-3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
