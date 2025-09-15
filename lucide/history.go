package lucide

import x "github.com/plainkit/blox"

// History creates a History Lucide icon.
func History(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-history", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 12a9 9 0 1 0 9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"))),
		x.Child(x.Path(x.D("M3 3v5h5"))),
		x.Child(x.Path(x.D("M12 7v5l4 2"))),
	)
	return x.Svg(svgArgs...)
}
