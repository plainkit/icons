package lucide

import x "github.com/plainkit/blox"

// GitPullRequestArrow creates a Git Pull Request Arrow Lucide icon.
func GitPullRequestArrow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-pull-request-arrow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("5"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M5 9v12"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("18"), x.R("3"))),
		x.Child(x.Path(x.D("m15 9-3-3 3-3"))),
		x.Child(x.Path(x.D("M12 6h5a2 2 0 0 1 2 2v7"))),
	)
	return x.Svg(svgArgs...)
}
