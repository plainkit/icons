package lucide

import x "github.com/plainkit/html"

// Users creates a Users Lucide icon.
func Users(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-users", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"))),
		x.Child(x.Path(x.D("M16 3.128a4 4 0 0 1 0 7.744"))),
		x.Child(x.Path(x.D("M22 21v-2a4 4 0 0 0-3-3.87"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("7"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
