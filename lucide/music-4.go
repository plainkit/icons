package lucide

import x "github.com/plainkit/blox"

// Music4 creates a Music 4 Lucide icon.
func Music4(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-music-4", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 18V5l12-2v13"))),
		x.Child(x.Path(x.D("m9 9 12-2"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("18"), x.R("3"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("16"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
