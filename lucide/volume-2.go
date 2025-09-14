package lucide

import x "github.com/bloxui/blox"

// Volume2 creates a Volume 2 Lucide icon.
func Volume2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-volume-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 4.702a.705.705 0 0 0-1.203-.498L6.413 7.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298z"))),
		x.Child(x.Path(x.D("M16 9a5 5 0 0 1 0 6"))),
		x.Child(x.Path(x.D("M19.364 18.364a9 9 0 0 0 0-12.728"))),
	)
	return x.Svg(svgArgs...)
}
