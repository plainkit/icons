package lucide

import x "github.com/plainkit/blox"

// Angry creates a Angry Lucide icon.
func Angry(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-angry", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M16 16s-1.5-2-4-2-4 2-4 2"))),
		x.Child(x.Path(x.D("M7.5 8 10 9"))),
		x.Child(x.Path(x.D("m14 9 2.5-1"))),
		x.Child(x.Path(x.D("M9 10h.01"))),
		x.Child(x.Path(x.D("M15 10h.01"))),
	)
	return x.Svg(svgArgs...)
}
