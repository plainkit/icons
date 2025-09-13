package lucide

import x "github.com/bloxui/blox"

// OctagonX creates a Octagon X Lucide icon.
func OctagonX(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-octagon-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 9-6 6"))),
		x.Child(x.Path(x.D("M2.586 16.726A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2h6.624a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586z"))),
		x.Child(x.Path(x.D("m9 9 6 6"))),
	)
	return x.Svg(svgArgs...)
}
