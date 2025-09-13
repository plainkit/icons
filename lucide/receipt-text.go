package lucide

import x "github.com/bloxui/blox"

// ReceiptText creates a Receipt Text Lucide icon.
func ReceiptText(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-receipt-text", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("M14 8H8"))),
		x.Child(x.Path(x.D("M16 12H8"))),
		x.Child(x.Path(x.D("M13 16H8"))),
	)
	return x.Svg(svgArgs...)
}
