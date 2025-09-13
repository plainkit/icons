package lucide

import x "github.com/bloxui/blox"

// Volume1 creates a Volume 1 Lucide icon.
func Volume1(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-volume-1", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 4.702a.705.705 0 0 0-1.203-.498L6.413 7.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298z"))),
		x.Child(x.Path(x.D("M16 9a5 5 0 0 1 0 6"))),
	)
	return x.Svg(svgArgs...)
}
