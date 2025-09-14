package lucide

import x "github.com/bloxui/blox"

// ArrowDownToDot creates a Arrow Down To Dot Lucide icon.
func ArrowDownToDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-to-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v14"))),
		x.Child(x.Path(x.D("m19 9-7 7-7-7"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("21"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
