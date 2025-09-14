package lucide

import x "github.com/bloxui/blox"

// Power creates a Power Lucide icon.
func Power(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-power", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v10"))),
		x.Child(x.Path(x.D("M18.4 6.6a9 9 0 1 1-12.77.04"))),
	)
	return x.Svg(svgArgs...)
}
