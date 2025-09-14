package lucide

import x "github.com/bloxui/blox"

// Wheat creates a Wheat Lucide icon.
func Wheat(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wheat", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 22 16 8"))),
		x.Child(x.Path(x.D("M3.47 12.53 5 11l1.53 1.53a3.5 3.5 0 0 1 0 4.94L5 19l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Z"))),
		x.Child(x.Path(x.D("M7.47 8.53 9 7l1.53 1.53a3.5 3.5 0 0 1 0 4.94L9 15l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Z"))),
		x.Child(x.Path(x.D("M11.47 4.53 13 3l1.53 1.53a3.5 3.5 0 0 1 0 4.94L13 11l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Z"))),
		x.Child(x.Path(x.D("M20 2h2v2a4 4 0 0 1-4 4h-2V6a4 4 0 0 1 4-4Z"))),
		x.Child(x.Path(x.D("M11.47 17.47 13 19l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L5 19l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"))),
		x.Child(x.Path(x.D("M15.47 13.47 17 15l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L9 15l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"))),
		x.Child(x.Path(x.D("M19.47 9.47 21 11l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L13 11l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"))),
	)
	return x.Svg(svgArgs...)
}
