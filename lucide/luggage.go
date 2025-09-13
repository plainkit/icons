package lucide

import x "github.com/bloxui/blox"

// Luggage creates a Luggage Lucide icon.
func Luggage(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-luggage", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 20a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M8 18V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v14"))),
		x.Child(x.Path(x.D("M10 20h4"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("20"), x.R("2"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("20"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
