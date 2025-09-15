package lucide

import x "github.com/plainkit/html"

// Scaling creates a Scaling Lucide icon.
func Scaling(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scaling", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"))),
		x.Child(x.Path(x.D("M14 15H9v-5"))),
		x.Child(x.Path(x.D("M16 3h5v5"))),
		x.Child(x.Path(x.D("M21 3 9 15"))),
	)
	return x.Svg(svgArgs...)
}
