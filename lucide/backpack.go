package lucide

import x "github.com/plainkit/blox"

// Backpack creates a Backpack Lucide icon.
func Backpack(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-backpack", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 10a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"))),
		x.Child(x.Path(x.D("M8 10h8"))),
		x.Child(x.Path(x.D("M8 18h8"))),
		x.Child(x.Path(x.D("M8 22v-6a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v6"))),
		x.Child(x.Path(x.D("M9 6V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2"))),
	)
	return x.Svg(svgArgs...)
}
