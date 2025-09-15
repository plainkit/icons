package lucide

import x "github.com/plainkit/blox"

// Split creates a Split Lucide icon.
func Split(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-split", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 3h5v5"))),
		x.Child(x.Path(x.D("M8 3H3v5"))),
		x.Child(x.Path(x.D("M12 22v-8.3a4 4 0 0 0-1.172-2.872L3 3"))),
		x.Child(x.Path(x.D("m15 9 6-6"))),
	)
	return x.Svg(svgArgs...)
}
