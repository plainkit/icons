package lucide

import x "github.com/bloxui/blox"

// ListChecks creates a List Checks Lucide icon.
func ListChecks(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-list-checks", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 5h8"))),
		x.Child(x.Path(x.D("M13 12h8"))),
		x.Child(x.Path(x.D("M13 19h8"))),
		x.Child(x.Path(x.D("m3 17 2 2 4-4"))),
		x.Child(x.Path(x.D("m3 7 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
