package lucide

import x "github.com/plainkit/blox"

// ReceiptSwissFranc creates a Receipt Swiss Franc Lucide icon.
func ReceiptSwissFranc(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-receipt-swiss-franc", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("M10 17V7h5"))),
		x.Child(x.Path(x.D("M10 11h4"))),
		x.Child(x.Path(x.D("M8 15h5"))),
	)
	return x.Svg(svgArgs...)
}
