package lucide

import x "github.com/bloxui/blox"

// GripHorizontal creates a Grip Horizontal Lucide icon.
func GripHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-grip-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("9"), x.R("1"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("9"), x.R("1"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("9"), x.R("1"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("15"), x.R("1"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("15"), x.R("1"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("15"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
