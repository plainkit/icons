package lucide

import x "github.com/bloxui/blox"

// Music creates a Music Lucide icon.
func Music(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-music", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 18V5l12-2v13"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("18"), x.R("3"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("16"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
