package lucide

import x "github.com/plainkit/blox"

// Axis3d creates a Axis 3d Lucide icon.
func Axis3d(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-axis-3d", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.5 10.5 15 9"))),
		x.Child(x.Path(x.D("M4 4v15a1 1 0 0 0 1 1h15"))),
		x.Child(x.Path(x.D("M4.293 19.707 6 18"))),
		x.Child(x.Path(x.D("m9 15 1.5-1.5"))),
	)
	return x.Svg(svgArgs...)
}
