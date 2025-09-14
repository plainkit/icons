package lucide

import x "github.com/bloxui/blox"

// Clapperboard creates a Clapperboard Lucide icon.
func Clapperboard(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clapperboard", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20.2 6 3 11l-.9-2.4c-.3-1.1.3-2.2 1.3-2.5l13.5-4c1.1-.3 2.2.3 2.5 1.3Z"))),
		x.Child(x.Path(x.D("m6.2 5.3 3.1 3.9"))),
		x.Child(x.Path(x.D("m12.4 3.4 3.1 4"))),
		x.Child(x.Path(x.D("M3 11h18v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"))),
	)
	return x.Svg(svgArgs...)
}
