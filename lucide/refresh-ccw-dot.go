package lucide

import x "github.com/plainkit/blox"

// RefreshCcwDot creates a Refresh Ccw Dot Lucide icon.
func RefreshCcwDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-refresh-ccw-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"))),
		x.Child(x.Path(x.D("M3 3v5h5"))),
		x.Child(x.Path(x.D("M3 12a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"))),
		x.Child(x.Path(x.D("M16 16h5v5"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
