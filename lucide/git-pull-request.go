package lucide

import (
	html "github.com/plainkit/html"
)

// GitPullRequest creates a Git Pull Request Lucide icon.
func GitPullRequest(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-pull-request", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3")),
		html.SvgCircle(html.ACx("6"), html.ACy("6"), html.AR("3")),
		html.SvgPath(html.AD("M13 6h3a2 2 0 0 1 2 2v7")),
		html.SvgLine(html.AX1("6"), html.AX2("6"), html.AY1("9"), html.AY2("21")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
