package lucide

import x "github.com/plainkit/html"

// SwissFranc creates a Swiss Franc Lucide icon.
func SwissFranc(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-swiss-franc", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 21V3h8"))),
		x.Child(x.Path(x.D("M6 16h9"))),
		x.Child(x.Path(x.D("M10 9.5h7"))),
	)
	return x.Svg(svgArgs...)
}
