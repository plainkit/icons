package lucide

import x "github.com/bloxui/blox"

// TextAlignCenter creates a Text Align Center Lucide icon.
func TextAlignCenter(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-text-align-center", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 5H3"))),
		x.Child(x.Path(x.D("M17 12H7"))),
		x.Child(x.Path(x.D("M19 19H5"))),
	)
	return x.Svg(svgArgs...)
}
