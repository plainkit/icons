package lucide

import x "github.com/bloxui/blox"

// ReceiptCent creates a Receipt Cent Lucide icon.
func ReceiptCent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-receipt-cent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("M12 6.5v11"))),
		x.Child(x.Path(x.D("M15 9.4a4 4 0 1 0 0 5.2"))),
	)
	return x.Svg(svgArgs...)
}
