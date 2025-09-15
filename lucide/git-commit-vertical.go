package lucide

import x "github.com/plainkit/html"

// GitCommitVertical creates a Git Commit Vertical Lucide icon.
func GitCommitVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-commit-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v6"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
		x.Child(x.Path(x.D("M12 15v6"))),
	)
	return x.Svg(svgArgs...)
}
