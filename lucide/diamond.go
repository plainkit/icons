package lucide

import x "github.com/bloxui/blox"

// Diamond creates a Diamond Lucide icon.
func Diamond(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-diamond", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41l-7.59-7.59a2.41 2.41 0 0 0-3.41 0Z"))),
	)
	return x.Svg(svgArgs...)
}
