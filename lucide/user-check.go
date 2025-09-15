package lucide

import x "github.com/plainkit/blox"

// UserCheck creates a User Check Lucide icon.
func UserCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 11 2 2 4-4"))),
		x.Child(x.Path(x.D("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("7"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
