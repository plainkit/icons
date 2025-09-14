package lucide

import x "github.com/bloxui/blox"

// Shuffle creates a Shuffle Lucide icon.
func Shuffle(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shuffle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m18 14 4 4-4 4"))),
		x.Child(x.Path(x.D("m18 2 4 4-4 4"))),
		x.Child(x.Path(x.D("M2 18h1.973a4 4 0 0 0 3.3-1.7l5.454-8.6a4 4 0 0 1 3.3-1.7H22"))),
		x.Child(x.Path(x.D("M2 6h1.972a4 4 0 0 1 3.6 2.2"))),
		x.Child(x.Path(x.D("M22 18h-6.041a4 4 0 0 1-3.3-1.8l-.359-.45"))),
	)
	return x.Svg(svgArgs...)
}
