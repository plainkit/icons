package lucide

import x "github.com/bloxui/blox"

// Castle creates a Castle Lucide icon.
func Castle(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-castle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 5V3"))),
		x.Child(x.Path(x.D("M14 5V3"))),
		x.Child(x.Path(x.D("M15 21v-3a3 3 0 0 0-6 0v3"))),
		x.Child(x.Path(x.D("M18 3v8"))),
		x.Child(x.Path(x.D("M18 5H6"))),
		x.Child(x.Path(x.D("M22 11H2"))),
		x.Child(x.Path(x.D("M22 9v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9"))),
		x.Child(x.Path(x.D("M6 3v8"))),
	)
	return x.Svg(svgArgs...)
}
