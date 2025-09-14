package lucide

import x "github.com/bloxui/blox"

// OctagonMinus creates a Octagon Minus Lucide icon.
func OctagonMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-octagon-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.586 16.726A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2h6.624a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586z"))),
		x.Child(x.Path(x.D("M8 12h8"))),
	)
	return x.Svg(svgArgs...)
}
