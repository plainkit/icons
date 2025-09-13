package lucide

import x "github.com/bloxui/blox"

// Music2 creates a Music 2 Lucide icon.
func Music2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-music-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("8"), x.Cy("18"), x.R("4"))),
		x.Child(x.Path(x.D("M12 18V2l7 4"))),
	)
	return x.Svg(svgArgs...)
}
