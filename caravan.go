package lucide

import x "github.com/bloxui/blox"

// Caravan creates a Caravan Lucide icon.
func Caravan(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-caravan", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 19V9a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v8a2 2 0 0 0 2 2h2"))),
		x.Child(x.Path(x.D("M2 9h3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H2"))),
		x.Child(x.Path(x.D("M22 17v1a1 1 0 0 1-1 1H10v-9a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v9"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("19"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
