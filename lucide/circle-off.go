package lucide

import x "github.com/plainkit/blox"

// CircleOff creates a Circle Off Lucide icon.
func CircleOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M8.35 2.69A10 10 0 0 1 21.3 15.65"))),
		x.Child(x.Path(x.D("M19.08 19.08A10 10 0 1 1 4.92 4.92"))),
	)
	return x.Svg(svgArgs...)
}
