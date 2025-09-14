package lucide

import x "github.com/bloxui/blox"

// TextAlignStart creates a Text Align Start Lucide icon.
func TextAlignStart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-align-start", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 5H3"))),
		x.Child(x.Path(x.D("M15 12H3"))),
		x.Child(x.Path(x.D("M17 19H3"))),
	)
	return x.Svg(svgArgs...)
}
