package lucide

import x "github.com/plainkit/blox"

// MapMinus creates a Map Minus Lucide icon.
func MapMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m11 19-1.106-.552a2 2 0 0 0-1.788 0l-3.659 1.83A1 1 0 0 1 3 19.381V6.618a1 1 0 0 1 .553-.894l4.553-2.277a2 2 0 0 1 1.788 0l4.212 2.106a2 2 0 0 0 1.788 0l3.659-1.83A1 1 0 0 1 21 4.619V14"))),
		x.Child(x.Path(x.D("M15 5.764V14"))),
		x.Child(x.Path(x.D("M21 18h-6"))),
		x.Child(x.Path(x.D("M9 3.236v15"))),
	)
	return x.Svg(svgArgs...)
}
