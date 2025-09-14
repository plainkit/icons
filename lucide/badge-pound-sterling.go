package lucide

import x "github.com/bloxui/blox"

// BadgePoundSterling creates a Badge Pound Sterling Lucide icon.
func BadgePoundSterling(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-badge-pound-sterling", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("M8 12h4"))),
		x.Child(x.Path(x.D("M10 16V9.5a2.5 2.5 0 0 1 5 0"))),
		x.Child(x.Path(x.D("M8 16h7"))),
	)
	return x.Svg(svgArgs...)
}
