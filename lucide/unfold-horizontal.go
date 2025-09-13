package lucide

import x "github.com/bloxui/blox"

// UnfoldHorizontal creates a Unfold Horizontal Lucide icon.
func UnfoldHorizontal(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-unfold-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 12h6"))),
		x.Child(x.Path(x.D("M8 12H2"))),
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("M12 8v2"))),
		x.Child(x.Path(x.D("M12 14v2"))),
		x.Child(x.Path(x.D("M12 20v2"))),
		x.Child(x.Path(x.D("m19 15 3-3-3-3"))),
		x.Child(x.Path(x.D("m5 9-3 3 3 3"))),
	)
	return x.Svg(svgArgs...)
}
