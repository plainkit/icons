package lucide

import x "github.com/bloxui/blox"

// PawPrint creates a Paw Print Lucide icon.
func PawPrint(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-paw-print", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("11"), x.Cy("4"), x.R("2"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("8"), x.R("2"))),
		x.Child(x.Circle(x.Cx("20"), x.Cy("16"), x.R("2"))),
		x.Child(x.Path(x.D("M9 10a5 5 0 0 1 5 5v3.5a3.5 3.5 0 0 1-6.84 1.045Q6.52 17.48 4.46 16.84A3.5 3.5 0 0 1 5.5 10Z"))),
	)
	return x.Svg(svgArgs...)
}
