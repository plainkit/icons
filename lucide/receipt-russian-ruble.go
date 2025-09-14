package lucide

import x "github.com/bloxui/blox"

// ReceiptRussianRuble creates a Receipt Russian Ruble Lucide icon.
func ReceiptRussianRuble(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-receipt-russian-ruble", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("M8 15h5"))),
		x.Child(x.Path(x.D("M8 11h5a2 2 0 1 0 0-4h-3v10"))),
	)
	return x.Svg(svgArgs...)
}
