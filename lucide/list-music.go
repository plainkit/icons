package lucide

import x "github.com/bloxui/blox"

// ListMusic creates a List Music Lucide icon.
func ListMusic(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-list-music", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5H3"))),
		x.Child(x.Path(x.D("M11 12H3"))),
		x.Child(x.Path(x.D("M11 19H3"))),
		x.Child(x.Path(x.D("M21 16V5"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("16"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
