package lucide

import x "github.com/plainkit/blox"

// Bike creates a Bike Lucide icon.
func Bike(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bike", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("18.5"), x.Cy("17.5"), x.R("3.5"))),
		x.Child(x.Circle(x.Cx("5.5"), x.Cy("17.5"), x.R("3.5"))),
		x.Child(x.Circle(x.Cx("15"), x.Cy("5"), x.R("1"))),
		x.Child(x.Path(x.D("M12 17.5V14l-3-3 4-3 2 3h2"))),
	)
	return x.Svg(svgArgs...)
}
