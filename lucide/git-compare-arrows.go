package lucide

import x "github.com/bloxui/blox"

// GitCompareArrows creates a Git Compare Arrows Lucide icon.
func GitCompareArrows(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-git-compare-arrows", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("5"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M12 6h5a2 2 0 0 1 2 2v7"))),
		x.Child(x.Path(x.D("m15 9-3-3 3-3"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("18"), x.R("3"))),
		x.Child(x.Path(x.D("M12 18H7a2 2 0 0 1-2-2V9"))),
		x.Child(x.Path(x.D("m9 15 3 3-3 3"))),
	)
	return x.Svg(svgArgs...)
}
