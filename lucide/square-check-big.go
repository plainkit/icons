package lucide

import x "github.com/bloxui/blox"

// SquareCheckBig creates a Square Check Big Lucide icon.
func SquareCheckBig(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-check-big", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 10.656V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h12.344"))),
		x.Child(x.Path(x.D("m9 11 3 3L22 4"))),
	)
	return x.Svg(svgArgs...)
}
