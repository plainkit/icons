package lucide

import x "github.com/bloxui/blox"

// UserX creates a User X Lucide icon.
func UserX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("7"), x.R("4"))),
		x.Child(x.Line(x.X1("17"), x.X2("22"), x.Y1("8"), x.Y2("13"))),
		x.Child(x.Line(x.X1("22"), x.X2("17"), x.Y1("8"), x.Y2("13"))),
	)
	return x.Svg(svgArgs...)
}
