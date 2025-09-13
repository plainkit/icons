package lucide

import x "github.com/bloxui/blox"

// RefreshCcw creates a Refresh Ccw Lucide icon.
func RefreshCcw(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-refresh-ccw", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"))),
		x.Child(x.Path(x.D("M3 3v5h5"))),
		x.Child(x.Path(x.D("M3 12a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"))),
		x.Child(x.Path(x.D("M16 16h5v5"))),
	)
	return x.Svg(svgArgs...)
}
