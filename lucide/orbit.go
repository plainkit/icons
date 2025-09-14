package lucide

import x "github.com/bloxui/blox"

// Orbit creates a Orbit Lucide icon.
func Orbit(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-orbit", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20.341 6.484A10 10 0 0 1 10.266 21.85"))),
		x.Child(x.Path(x.D("M3.659 17.516A10 10 0 0 1 13.74 2.152"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("5"), x.R("2"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("19"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
