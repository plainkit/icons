package lucide

import x "github.com/bloxui/blox"

// Cake creates a Cake Lucide icon.
func Cake(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cake", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 21v-8a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v8"))),
		x.Child(x.Path(x.D("M4 16s.5-1 2-1 2.5 2 4 2 2.5-2 4-2 2.5 2 4 2 2-1 2-1"))),
		x.Child(x.Path(x.D("M2 21h20"))),
		x.Child(x.Path(x.D("M7 8v3"))),
		x.Child(x.Path(x.D("M12 8v3"))),
		x.Child(x.Path(x.D("M17 8v3"))),
		x.Child(x.Path(x.D("M7 4h.01"))),
		x.Child(x.Path(x.D("M12 4h.01"))),
		x.Child(x.Path(x.D("M17 4h.01"))),
	)
	return x.Svg(svgArgs...)
}
