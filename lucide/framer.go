package lucide

import x "github.com/plainkit/blox"

// Framer creates a Framer Lucide icon.
func Framer(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-framer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 16V9h14V2H5l14 14h-7m-7 0 7 7v-7m-7 0h7"))),
	)
	return x.Svg(svgArgs...)
}
