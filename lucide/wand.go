package lucide

import x "github.com/plainkit/html"

// Wand creates a Wand Lucide icon.
func Wand(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wand", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 4V2"))),
		x.Child(x.Path(x.D("M15 16v-2"))),
		x.Child(x.Path(x.D("M8 9h2"))),
		x.Child(x.Path(x.D("M20 9h2"))),
		x.Child(x.Path(x.D("M17.8 11.8 19 13"))),
		x.Child(x.Path(x.D("M15 9h.01"))),
		x.Child(x.Path(x.D("M17.8 6.2 19 5"))),
		x.Child(x.Path(x.D("m3 21 9-9"))),
		x.Child(x.Path(x.D("M12.2 6.2 11 5"))),
	)
	return x.Svg(svgArgs...)
}
