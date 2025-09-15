package lucide

import x "github.com/plainkit/html"

// GitCompare creates a Git Compare Lucide icon.
func GitCompare(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-compare", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M13 6h3a2 2 0 0 1 2 2v7"))),
		x.Child(x.Path(x.D("M11 18H8a2 2 0 0 1-2-2V9"))),
	)
	return x.Svg(svgArgs...)
}
