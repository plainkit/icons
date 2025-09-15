package lucide

import x "github.com/plainkit/html"

// RotateCcw creates a Rotate Ccw Lucide icon.
func RotateCcw(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rotate-ccw", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 12a9 9 0 1 0 9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"))),
		x.Child(x.Path(x.D("M3 3v5h5"))),
	)
	return x.Svg(svgArgs...)
}
