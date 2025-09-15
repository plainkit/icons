package lucide

import x "github.com/plainkit/blox"

// Merge creates a Merge Lucide icon.
func Merge(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-merge", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m8 6 4-4 4 4"))),
		x.Child(x.Path(x.D("M12 2v10.3a4 4 0 0 1-1.172 2.872L4 22"))),
		x.Child(x.Path(x.D("m20 22-5-5"))),
	)
	return x.Svg(svgArgs...)
}
