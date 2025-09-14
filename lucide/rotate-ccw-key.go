package lucide

import x "github.com/bloxui/blox"

// RotateCcwKey creates a Rotate Ccw Key Lucide icon.
func RotateCcwKey(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rotate-ccw-key", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14.5 9.5 1 1"))),
		x.Child(x.Path(x.D("m15.5 8.5-4 4"))),
		x.Child(x.Path(x.D("M3 12a9 9 0 1 0 9-9 9.74 9.74 0 0 0-6.74 2.74L3 8"))),
		x.Child(x.Path(x.D("M3 3v5h5"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("14"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
