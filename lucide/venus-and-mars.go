package lucide

import x "github.com/bloxui/blox"

// VenusAndMars creates a Venus And Mars Lucide icon.
func VenusAndMars(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-venus-and-mars", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 20h4"))),
		x.Child(x.Path(x.D("M12 16v6"))),
		x.Child(x.Path(x.D("M17 2h4v4"))),
		x.Child(x.Path(x.D("m21 2-5.46 5.46"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("11"), x.R("5"))),
	)
	return x.Svg(svgArgs...)
}
