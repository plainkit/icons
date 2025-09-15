package lucide

import x "github.com/plainkit/blox"

// PilcrowLeft creates a Pilcrow Left Lucide icon.
func PilcrowLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pilcrow-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 3v11"))),
		x.Child(x.Path(x.D("M14 9h-3a3 3 0 0 1 0-6h9"))),
		x.Child(x.Path(x.D("M18 3v11"))),
		x.Child(x.Path(x.D("M22 18H2l4-4"))),
		x.Child(x.Path(x.D("m6 22-4-4"))),
	)
	return x.Svg(svgArgs...)
}
