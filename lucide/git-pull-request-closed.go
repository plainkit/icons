package lucide

import (
	html "github.com/plainkit/html"
)

// GitPullRequestClosed creates a Git Pull Request Closed Lucide icon.
func GitPullRequestClosed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-pull-request-closed", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("6"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M6 9v12"))),
		html.Child(html.SvgPath(html.AD("m21 3-6 6"))),
		html.Child(html.SvgPath(html.AD("m21 9-6-6"))),
		html.Child(html.SvgPath(html.AD("M18 11.5V15"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
