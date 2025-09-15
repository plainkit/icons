package lucide

import x "github.com/plainkit/html"

// PoundSterling creates a Pound Sterling Lucide icon.
func PoundSterling(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pound-sterling", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 7c0-5.333-8-5.333-8 0"))),
		x.Child(x.Path(x.D("M10 7v14"))),
		x.Child(x.Path(x.D("M6 21h12"))),
		x.Child(x.Path(x.D("M6 13h10"))),
	)
	return x.Svg(svgArgs...)
}
