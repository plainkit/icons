package lucide

import x "github.com/bloxui/blox"

// Radius creates a Radius Lucide icon.
func Radius(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-radius", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20.34 17.52a10 10 0 1 0-2.82 2.82"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("19"), x.R("2"))),
		x.Child(x.Path(x.D("m13.41 13.41 4.18 4.18"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
