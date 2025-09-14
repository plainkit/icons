package lucide

import x "github.com/bloxui/blox"

// Globe creates a Globe Lucide icon.
func Globe(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-globe", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20"))),
		x.Child(x.Path(x.D("M2 12h20"))),
	)
	return x.Svg(svgArgs...)
}
