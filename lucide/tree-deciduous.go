package lucide

import x "github.com/plainkit/html"

// TreeDeciduous creates a Tree Deciduous Lucide icon.
func TreeDeciduous(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tree-deciduous", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 19a4 4 0 0 1-2.24-7.32A3.5 3.5 0 0 1 9 6.03V6a3 3 0 1 1 6 0v.04a3.5 3.5 0 0 1 3.24 5.65A4 4 0 0 1 16 19Z"))),
		x.Child(x.Path(x.D("M12 19v3"))),
	)
	return x.Svg(svgArgs...)
}
