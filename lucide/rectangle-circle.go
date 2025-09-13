package lucide

import x "github.com/bloxui/blox"

// RectangleCircle creates a Rectangle Circle Lucide icon.
func RectangleCircle(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-rectangle-circle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 4v16H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1z"))),
		x.Child(x.Circle(x.Cx("14"), x.Cy("12"), x.R("8"))),
	)
	return x.Svg(svgArgs...)
}
