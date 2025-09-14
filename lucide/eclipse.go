package lucide

import x "github.com/bloxui/blox"

// Eclipse creates a Eclipse Lucide icon.
func Eclipse(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-eclipse", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M12 2a7 7 0 1 0 10 10"))),
	)
	return x.Svg(svgArgs...)
}
