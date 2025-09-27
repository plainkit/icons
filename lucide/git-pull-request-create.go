package lucide

import (
	html "github.com/plainkit/html"
)

// GitPullRequestCreate creates a Git Pull Request Create Lucide icon.
func GitPullRequestCreate(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-pull-request-create", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("6"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M6 9v12"))),
		html.Child(html.SvgPath(html.AD("M13 6h3a2 2 0 0 1 2 2v3"))),
		html.Child(html.SvgPath(html.AD("M18 15v6"))),
		html.Child(html.SvgPath(html.AD("M21 18h-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
