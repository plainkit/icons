package lucide

import x "github.com/bloxui/blox"

// ReceiptPoundSterling creates a Receipt Pound Sterling Lucide icon.
func ReceiptPoundSterling(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-receipt-pound-sterling", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("M8 13h5"))),
		x.Child(x.Path(x.D("M10 17V9.5a2.5 2.5 0 0 1 5 0"))),
		x.Child(x.Path(x.D("M8 17h7"))),
	)
	return x.Svg(svgArgs...)
}
