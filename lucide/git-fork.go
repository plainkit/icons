package lucide

import x "github.com/plainkit/blox"

// GitFork creates a Git Fork Lucide icon.
func GitFork(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-fork", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("18"), x.R("3"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("6"), x.R("3"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M18 9v2c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V9"))),
		x.Child(x.Path(x.D("M12 12v3"))),
	)
	return x.Svg(svgArgs...)
}
