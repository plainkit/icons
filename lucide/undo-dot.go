package lucide

import x "github.com/plainkit/blox"

// UndoDot creates a Undo Dot Lucide icon.
func UndoDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-undo-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 17a9 9 0 0 0-15-6.7L3 13"))),
		x.Child(x.Path(x.D("M3 7v6h6"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("17"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
