package lucide

import x "github.com/bloxui/blox"

// UnfoldVertical creates a Unfold Vertical Lucide icon.
func UnfoldVertical(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-unfold-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 22v-6"))),
		x.Child(x.Path(x.D("M12 8V2"))),
		x.Child(x.Path(x.D("M4 12H2"))),
		x.Child(x.Path(x.D("M10 12H8"))),
		x.Child(x.Path(x.D("M16 12h-2"))),
		x.Child(x.Path(x.D("M22 12h-2"))),
		x.Child(x.Path(x.D("m15 19-3 3-3-3"))),
		x.Child(x.Path(x.D("m15 5-3-3-3 3"))),
	)
	return x.Svg(svgArgs...)
}
