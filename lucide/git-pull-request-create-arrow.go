package lucide

import x "github.com/bloxui/blox"

// GitPullRequestCreateArrow creates a Git Pull Request Create Arrow Lucide icon.
func GitPullRequestCreateArrow(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-git-pull-request-create-arrow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("5"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M5 9v12"))),
		x.Child(x.Path(x.D("m15 9-3-3 3-3"))),
		x.Child(x.Path(x.D("M12 6h5a2 2 0 0 1 2 2v3"))),
		x.Child(x.Path(x.D("M19 15v6"))),
		x.Child(x.Path(x.D("M22 18h-6"))),
	)
	return x.Svg(svgArgs...)
}
