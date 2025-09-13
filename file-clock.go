package lucide

import x "github.com/bloxui/blox"

// FileClock creates a File Clock Lucide icon.
func FileClock(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-clock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M16 22h2a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3"))),
		x.Child(x.Path(x.D("M8 14v2.2l1.6 1"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("16"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}