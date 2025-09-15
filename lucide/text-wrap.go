package lucide

import x "github.com/plainkit/blox"

// TextWrap creates a Text Wrap Lucide icon.
func TextWrap(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-wrap", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 16-3 3 3 3"))),
		x.Child(x.Path(x.D("M3 12h14.5a1 1 0 0 1 0 7H13"))),
		x.Child(x.Path(x.D("M3 19h6"))),
		x.Child(x.Path(x.D("M3 5h18"))),
	)
	return x.Svg(svgArgs...)
}
