package lucide

import x "github.com/plainkit/html"

// TextAlignJustify creates a Text Align Justify Lucide icon.
func TextAlignJustify(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-align-justify", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 5h18"))),
		x.Child(x.Path(x.D("M3 12h18"))),
		x.Child(x.Path(x.D("M3 19h18"))),
	)
	return x.Svg(svgArgs...)
}
