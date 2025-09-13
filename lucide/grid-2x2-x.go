package lucide

import x "github.com/bloxui/blox"

// Grid2x2X creates a Grid 2x2 X Lucide icon.
func Grid2x2X(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-grid-2x2-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3v17a1 1 0 0 1-1 1H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6a1 1 0 0 1-1 1H3"))),
		x.Child(x.Path(x.D("m16 16 5 5"))),
		x.Child(x.Path(x.D("m16 21 5-5"))),
	)
	return x.Svg(svgArgs...)
}
