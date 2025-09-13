package lucide

import x "github.com/bloxui/blox"

// Move creates a Move Lucide icon.
func Move(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-move", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v20"))),
		x.Child(x.Path(x.D("m15 19-3 3-3-3"))),
		x.Child(x.Path(x.D("m19 9 3 3-3 3"))),
		x.Child(x.Path(x.D("M2 12h20"))),
		x.Child(x.Path(x.D("m5 9-3 3 3 3"))),
		x.Child(x.Path(x.D("m9 5 3-3 3 3"))),
	)
	return x.Svg(svgArgs...)
}
