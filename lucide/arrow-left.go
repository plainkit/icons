package lucide

import x "github.com/bloxui/blox"

// ArrowLeft creates a Arrow Left Lucide icon.
func ArrowLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 19-7-7 7-7"))),
		x.Child(x.Path(x.D("M19 12H5"))),
	)
	return x.Svg(svgArgs...)
}
