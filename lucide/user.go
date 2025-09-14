package lucide

import x "github.com/bloxui/blox"

// User creates a User Lucide icon.
func User(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("7"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
