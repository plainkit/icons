package lucide

import (
	html "github.com/plainkit/html"
)

// GitPullRequestArrow creates a Git Pull Request Arrow Lucide icon.
func GitPullRequestArrow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-pull-request-arrow", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("5"), html.ACy("6"), html.AR("3")),
		html.SvgPath(html.AD("M5 9v12")),
		html.SvgCircle(html.ACx("19"), html.ACy("18"), html.AR("3")),
		html.SvgPath(html.AD("m15 9-3-3 3-3")),
		html.SvgPath(html.AD("M12 6h5a2 2 0 0 1 2 2v7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
