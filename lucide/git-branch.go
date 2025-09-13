package lucide

import x "github.com/bloxui/blox"

// GitBranch creates a Git Branch Lucide icon.
func GitBranch(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-git-branch", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("6"), x.X2("6"), x.Y1("3"), x.Y2("15"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("6"), x.R("3"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("18"), x.R("3"))),
		x.Child(x.Path(x.D("M18 9a9 9 0 0 1-9 9"))),
	)
	return x.Svg(svgArgs...)
}
