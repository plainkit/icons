package lucide

import x "github.com/bloxui/blox"

// ReceiptEuro creates a Receipt Euro Lucide icon.
func ReceiptEuro(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-receipt-euro", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("M8 12h5"))),
		x.Child(x.Path(x.D("M16 9.5a4 4 0 1 0 0 5.2"))),
	)
	return x.Svg(svgArgs...)
}
