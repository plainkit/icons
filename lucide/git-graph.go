package lucide

import x "github.com/bloxui/blox"

// GitGraph creates a Git Graph Lucide icon.
func GitGraph(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-graph", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("5"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M5 9v6"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("18"), x.R("3"))),
		x.Child(x.Path(x.D("M12 3v18"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M16 15.7A9 9 0 0 0 19 9"))),
	)
	return x.Svg(svgArgs...)
}
