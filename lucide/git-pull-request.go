package lucide

import x "github.com/bloxui/blox"

// GitPullRequest creates a Git Pull Request Lucide icon.
func GitPullRequest(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-pull-request", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M13 6h3a2 2 0 0 1 2 2v7"))),
		x.Child(x.Line(x.X1("6"), x.X2("6"), x.Y1("9"), x.Y2("21"))),
	)
	return x.Svg(svgArgs...)
}
