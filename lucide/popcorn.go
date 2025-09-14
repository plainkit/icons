package lucide

import x "github.com/bloxui/blox"

// Popcorn creates a Popcorn Lucide icon.
func Popcorn(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-popcorn", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 8a2 2 0 0 0 0-4 2 2 0 0 0-4 0 2 2 0 0 0-4 0 2 2 0 0 0-4 0 2 2 0 0 0 0 4"))),
		x.Child(x.Path(x.D("M10 22 9 8"))),
		x.Child(x.Path(x.D("m14 22 1-14"))),
		x.Child(x.Path(x.D("M20 8c.5 0 .9.4.8 1l-2.6 12c-.1.5-.7 1-1.2 1H7c-.6 0-1.1-.4-1.2-1L3.2 9c-.1-.6.3-1 .8-1Z"))),
	)
	return x.Svg(svgArgs...)
}
