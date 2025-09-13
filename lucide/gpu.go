package lucide

import x "github.com/bloxui/blox"

// Gpu creates a Gpu Lucide icon.
func Gpu(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-gpu", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 21V3"))),
		x.Child(x.Path(x.D("M2 5h18a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2.26"))),
		x.Child(x.Path(x.D("M7 17v3a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-3"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("11"), x.R("2"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("11"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
