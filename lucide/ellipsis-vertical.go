package lucide

import x "github.com/bloxui/blox"

// EllipsisVertical creates a Ellipsis Vertical Lucide icon.
func EllipsisVertical(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-ellipsis-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("5"), x.R("1"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("19"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
