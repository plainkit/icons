package lucide

import x "github.com/bloxui/blox"

// Citrus creates a Citrus Lucide icon.
func Citrus(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-citrus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21.66 17.67a1.08 1.08 0 0 1-.04 1.6A12 12 0 0 1 4.73 2.38a1.1 1.1 0 0 1 1.61-.04z"))),
		x.Child(x.Path(x.D("M19.65 15.66A8 8 0 0 1 8.35 4.34"))),
		x.Child(x.Path(x.D("m14 10-5.5 5.5"))),
		x.Child(x.Path(x.D("M14 17.85V10H6.15"))),
	)
	return x.Svg(svgArgs...)
}
