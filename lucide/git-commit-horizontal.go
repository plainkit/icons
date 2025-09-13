package lucide

import x "github.com/bloxui/blox"

// GitCommitHorizontal creates a Git Commit Horizontal Lucide icon.
func GitCommitHorizontal(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-git-commit-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
		x.Child(x.Line(x.X1("3"), x.X2("9"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("15"), x.X2("21"), x.Y1("12"), x.Y2("12"))),
	)
	return x.Svg(svgArgs...)
}
