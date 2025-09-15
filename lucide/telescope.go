package lucide

import x "github.com/plainkit/blox"

// Telescope creates a Telescope Lucide icon.
func Telescope(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-telescope", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10.065 12.493-6.18 1.318a.934.934 0 0 1-1.108-.702l-.537-2.15a1.07 1.07 0 0 1 .691-1.265l13.504-4.44"))),
		x.Child(x.Path(x.D("m13.56 11.747 4.332-.924"))),
		x.Child(x.Path(x.D("m16 21-3.105-6.21"))),
		x.Child(x.Path(x.D("M16.485 5.94a2 2 0 0 1 1.455-2.425l1.09-.272a1 1 0 0 1 1.212.727l1.515 6.06a1 1 0 0 1-.727 1.213l-1.09.272a2 2 0 0 1-2.425-1.455z"))),
		x.Child(x.Path(x.D("m6.158 8.633 1.114 4.456"))),
		x.Child(x.Path(x.D("m8 21 3.105-6.21"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
