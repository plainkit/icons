package lucide

import x "github.com/plainkit/blox"

// Salad creates a Salad Lucide icon.
func Salad(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-salad", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 21h10"))),
		x.Child(x.Path(x.D("M12 21a9 9 0 0 0 9-9H3a9 9 0 0 0 9 9Z"))),
		x.Child(x.Path(x.D("M11.38 12a2.4 2.4 0 0 1-.4-4.77 2.4 2.4 0 0 1 3.2-2.77 2.4 2.4 0 0 1 3.47-.63 2.4 2.4 0 0 1 3.37 3.37 2.4 2.4 0 0 1-1.1 3.7 2.51 2.51 0 0 1 .03 1.1"))),
		x.Child(x.Path(x.D("m13 12 4-4"))),
		x.Child(x.Path(x.D("M10.9 7.25A3.99 3.99 0 0 0 4 10c0 .73.2 1.41.54 2"))),
	)
	return x.Svg(svgArgs...)
}
