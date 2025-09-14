package lucide

import x "github.com/bloxui/blox"

// Ruler creates a Ruler Lucide icon.
func Ruler(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ruler", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21.3 15.3a2.4 2.4 0 0 1 0 3.4l-2.6 2.6a2.4 2.4 0 0 1-3.4 0L2.7 8.7a2.41 2.41 0 0 1 0-3.4l2.6-2.6a2.41 2.41 0 0 1 3.4 0Z"))),
		x.Child(x.Path(x.D("m14.5 12.5 2-2"))),
		x.Child(x.Path(x.D("m11.5 9.5 2-2"))),
		x.Child(x.Path(x.D("m8.5 6.5 2-2"))),
		x.Child(x.Path(x.D("m17.5 15.5 2-2"))),
	)
	return x.Svg(svgArgs...)
}
