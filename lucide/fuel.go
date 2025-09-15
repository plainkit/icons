package lucide

import x "github.com/plainkit/html"

// Fuel creates a Fuel Lucide icon.
func Fuel(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-fuel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 13h2a2 2 0 0 1 2 2v2a2 2 0 0 0 4 0v-6.998a2 2 0 0 0-.59-1.42L18 5"))),
		x.Child(x.Path(x.D("M14 21V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v16"))),
		x.Child(x.Path(x.D("M2 21h13"))),
		x.Child(x.Path(x.D("M3 9h11"))),
	)
	return x.Svg(svgArgs...)
}
