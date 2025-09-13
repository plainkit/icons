package lucide

import x "github.com/bloxui/blox"

// Dribbble creates a Dribbble Lucide icon.
func Dribbble(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-dribbble", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M19.13 5.09C15.22 9.14 10 10.44 2.25 10.94"))),
		x.Child(x.Path(x.D("M21.75 12.84c-6.62-1.41-12.14 1-16.38 6.32"))),
		x.Child(x.Path(x.D("M8.56 2.75c4.37 6 6 9.42 8 17.72"))),
	)
	return x.Svg(svgArgs...)
}