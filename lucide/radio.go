package lucide

import x "github.com/plainkit/blox"

// Radio creates a Radio Lucide icon.
func Radio(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-radio", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16.247 7.761a6 6 0 0 1 0 8.478"))),
		x.Child(x.Path(x.D("M19.075 4.933a10 10 0 0 1 0 14.134"))),
		x.Child(x.Path(x.D("M4.925 19.067a10 10 0 0 1 0-14.134"))),
		x.Child(x.Path(x.D("M7.753 16.239a6 6 0 0 1 0-8.478"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
