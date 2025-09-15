package lucide

import x "github.com/plainkit/blox"

// GitPullRequestDraft creates a Git Pull Request Draft Lucide icon.
func GitPullRequestDraft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-pull-request-draft", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M18 6V5"))),
		x.Child(x.Path(x.D("M18 11v-1"))),
		x.Child(x.Line(x.X1("6"), x.X2("6"), x.Y1("9"), x.Y2("21"))),
	)
	return x.Svg(svgArgs...)
}
