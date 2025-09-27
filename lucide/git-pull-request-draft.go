package lucide

import (
	html "github.com/plainkit/html"
)

// GitPullRequestDraft creates a Git Pull Request Draft Lucide icon.
func GitPullRequestDraft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-pull-request-draft", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("6"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M18 6V5"))),
		html.Child(html.SvgPath(html.AD("M18 11v-1"))),
		html.Child(html.SvgLine(html.AX1("6"), html.AX2("6"), html.AY1("9"), html.AY2("21"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
