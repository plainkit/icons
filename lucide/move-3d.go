package lucide

import x "github.com/bloxui/blox"

// Move3d creates a Move 3d Lucide icon.
func Move3d(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-move-3d", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 3v16h16"))),
		x.Child(x.Path(x.D("m5 19 6-6"))),
		x.Child(x.Path(x.D("m2 6 3-3 3 3"))),
		x.Child(x.Path(x.D("m18 16 3 3-3 3"))),
	)
	return x.Svg(svgArgs...)
}
