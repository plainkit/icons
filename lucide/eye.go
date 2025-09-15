package lucide

import x "github.com/plainkit/blox"

// Eye creates a Eye Lucide icon.
func Eye(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-eye", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.062 12.348a1 1 0 0 1 0-.696 10.75 10.75 0 0 1 19.876 0 1 1 0 0 1 0 .696 10.75 10.75 0 0 1-19.876 0"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
