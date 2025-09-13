package lucide

import x "github.com/bloxui/blox"

// RefreshCw creates a Refresh Cw Lucide icon.
func RefreshCw(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-refresh-cw", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 12a9 9 0 0 1 9-9 9.75 9.75 0 0 1 6.74 2.74L21 8"))),
		x.Child(x.Path(x.D("M21 3v5h-5"))),
		x.Child(x.Path(x.D("M21 12a9 9 0 0 1-9 9 9.75 9.75 0 0 1-6.74-2.74L3 16"))),
		x.Child(x.Path(x.D("M8 16H3v5"))),
	)
	return x.Svg(svgArgs...)
}
