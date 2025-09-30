package lucide

import (
	html "github.com/plainkit/html"
)

// GitCommitHorizontal creates a Git Commit Horizontal Lucide icon.
func GitCommitHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-commit-horizontal", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
		html.SvgLine(html.AX1("3"), html.AX2("9"), html.AY1("12"), html.AY2("12")),
		html.SvgLine(html.AX1("15"), html.AX2("21"), html.AY1("12"), html.AY2("12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
