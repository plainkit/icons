package lucide

import x "github.com/plainkit/blox"

// MountainSnow creates a Mountain Snow Lucide icon.
func MountainSnow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mountain-snow", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m8 3 4 8 5-5 5 15H2L8 3z"))),
		x.Child(x.Path(x.D("M4.14 15.08c2.62-1.57 5.24-1.43 7.86.42 2.74 1.94 5.49 2 8.23.19"))),
	)
	return x.Svg(svgArgs...)
}
