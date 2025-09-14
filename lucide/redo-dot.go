package lucide

import x "github.com/bloxui/blox"

// RedoDot creates a Redo Dot Lucide icon.
func RedoDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-redo-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("17"), x.R("1"))),
		x.Child(x.Path(x.D("M21 7v6h-6"))),
		x.Child(x.Path(x.D("M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7"))),
	)
	return x.Svg(svgArgs...)
}
