package lucide

import x "github.com/bloxui/blox"

// UserSearch creates a User Search Lucide icon.
func UserSearch(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-user-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("10"), x.Cy("7"), x.R("4"))),
		x.Child(x.Path(x.D("M10.3 15H7a4 4 0 0 0-4 4v2"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("17"), x.R("3"))),
		x.Child(x.Path(x.D("m21 21-1.9-1.9"))),
	)
	return x.Svg(svgArgs...)
}
