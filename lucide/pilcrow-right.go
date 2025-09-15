package lucide

import x "github.com/plainkit/blox"

// PilcrowRight creates a Pilcrow Right Lucide icon.
func PilcrowRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pilcrow-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 3v11"))),
		x.Child(x.Path(x.D("M10 9H7a1 1 0 0 1 0-6h8"))),
		x.Child(x.Path(x.D("M14 3v11"))),
		x.Child(x.Path(x.D("m18 14 4 4H2"))),
		x.Child(x.Path(x.D("m22 18-4 4"))),
	)
	return x.Svg(svgArgs...)
}
