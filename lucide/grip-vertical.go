package lucide

import x "github.com/bloxui/blox"

// GripVertical creates a Grip Vertical Lucide icon.
func GripVertical(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-grip-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("9"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("5"), x.R("1"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("19"), x.R("1"))),
		x.Child(x.Circle(x.Cx("15"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("15"), x.Cy("5"), x.R("1"))),
		x.Child(x.Circle(x.Cx("15"), x.Cy("19"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
