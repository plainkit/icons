package lucide

import x "github.com/bloxui/blox"

// Forklift creates a Forklift Lucide icon.
func Forklift(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-forklift", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 12H5a2 2 0 0 0-2 2v5"))),
		x.Child(x.Circle(x.Cx("13"), x.Cy("19"), x.R("2"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("19"), x.R("2"))),
		x.Child(x.Path(x.D("M8 19h3m5-17v17h6M6 12V7c0-1.1.9-2 2-2h3l5 5"))),
	)
	return x.Svg(svgArgs...)
}
