package lucide

import x "github.com/bloxui/blox"

// Disc creates a Disc Lucide icon.
func Disc(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-disc", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
