package lucide

import x "github.com/bloxui/blox"

// TextAlignEnd creates a Text Align End Lucide icon.
func TextAlignEnd(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-text-align-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 5H3"))),
		x.Child(x.Path(x.D("M21 12H9"))),
		x.Child(x.Path(x.D("M21 19H7"))),
	)
	return x.Svg(svgArgs...)
}
