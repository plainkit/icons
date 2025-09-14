package lucide

import x "github.com/bloxui/blox"

// Forward creates a Forward Lucide icon.
func Forward(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-forward", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 17 5-5-5-5"))),
		x.Child(x.Path(x.D("M4 18v-2a4 4 0 0 1 4-4h12"))),
	)
	return x.Svg(svgArgs...)
}
