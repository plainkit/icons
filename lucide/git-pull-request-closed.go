package lucide

import x "github.com/bloxui/blox"

// GitPullRequestClosed creates a Git Pull Request Closed Lucide icon.
func GitPullRequestClosed(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-git-pull-request-closed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("6"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M6 9v12"))),
		x.Child(x.Path(x.D("m21 3-6 6"))),
		x.Child(x.Path(x.D("m21 9-6-6"))),
		x.Child(x.Path(x.D("M18 11.5V15"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
