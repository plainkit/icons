package lucide

import x "github.com/plainkit/html"

// GitPullRequestCreate creates a Git Pull Request Create Lucide icon.
func GitPullRequestCreate(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-pull-request-create", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("6"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M6 9v12"))),
		x.Child(x.Path(x.D("M13 6h3a2 2 0 0 1 2 2v3"))),
		x.Child(x.Path(x.D("M18 15v6"))),
		x.Child(x.Path(x.D("M21 18h-6"))),
	)
	return x.Svg(svgArgs...)
}
