package lucide

import (
	html "github.com/plainkit/html"
)

// GitBranch creates a Git Branch Lucide icon.
func GitBranch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-branch", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("6"), html.AX2("6"), html.AY1("3"), html.AY2("15")),
		html.SvgCircle(html.ACx("18"), html.ACy("6"), html.AR("3")),
		html.SvgCircle(html.ACx("6"), html.ACy("18"), html.AR("3")),
		html.SvgPath(html.AD("M18 9a9 9 0 0 1-9 9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
