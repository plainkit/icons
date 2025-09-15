package lucide

import x "github.com/plainkit/html"

// Map creates a Map Lucide icon.
func Map(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14.106 5.553a2 2 0 0 0 1.788 0l3.659-1.83A1 1 0 0 1 21 4.619v12.764a1 1 0 0 1-.553.894l-4.553 2.277a2 2 0 0 1-1.788 0l-4.212-2.106a2 2 0 0 0-1.788 0l-3.659 1.83A1 1 0 0 1 3 19.381V6.618a1 1 0 0 1 .553-.894l4.553-2.277a2 2 0 0 1 1.788 0z"))),
		x.Child(x.Path(x.D("M15 5.764v15"))),
		x.Child(x.Path(x.D("M9 3.236v15"))),
	)
	return x.Svg(svgArgs...)
}
