package lucide

import x "github.com/bloxui/blox"

// Blend creates a Blend Lucide icon.
func Blend(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-blend", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("9"), x.Cy("9"), x.R("7"))),
		x.Child(x.Circle(x.Cx("15"), x.Cy("15"), x.R("7"))),
	)
	return x.Svg(svgArgs...)
}
