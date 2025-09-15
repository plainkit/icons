package lucide

import x "github.com/plainkit/blox"

// Venus creates a Venus Lucide icon.
func Venus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-venus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 15v7"))),
		x.Child(x.Path(x.D("M9 19h6"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("9"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}
