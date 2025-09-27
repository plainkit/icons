package lucide

import (
	html "github.com/plainkit/html"
)

// GitPullRequestCreateArrow creates a Git Pull Request Create Arrow Lucide icon.
func GitPullRequestCreateArrow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-pull-request-create-arrow", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("5"), html.ACy("6"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M5 9v12"))),
		html.Child(html.SvgPath(html.AD("m15 9-3-3 3-3"))),
		html.Child(html.SvgPath(html.AD("M12 6h5a2 2 0 0 1 2 2v3"))),
		html.Child(html.SvgPath(html.AD("M19 15v6"))),
		html.Child(html.SvgPath(html.AD("M22 18h-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
