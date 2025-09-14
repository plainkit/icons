package lucide

import x "github.com/bloxui/blox"

// House creates a House Lucide icon.
func House(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-house", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 21v-8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v8"))),
		x.Child(x.Path(x.D("M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"))),
	)
	return x.Svg(svgArgs...)
}
