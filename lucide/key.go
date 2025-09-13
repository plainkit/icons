package lucide

import x "github.com/bloxui/blox"

// Key creates a Key Lucide icon.
func Key(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-key", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15.5 7.5 2.3 2.3a1 1 0 0 0 1.4 0l2.1-2.1a1 1 0 0 0 0-1.4L19 4"))),
		x.Child(x.Path(x.D("m21 2-9.6 9.6"))),
		x.Child(x.Circle(x.Cx("7.5"), x.Cy("15.5"), x.R("5.5"))),
	)
	return x.Svg(svgArgs...)
}
