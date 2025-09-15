package lucide

import x "github.com/plainkit/blox"

// ReceiptTurkishLira creates a Receipt Turkish Lira Lucide icon.
func ReceiptTurkishLira(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-receipt-turkish-lira", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 6.5v11a5.5 5.5 0 0 0 5.5-5.5"))),
		x.Child(x.Path(x.D("m14 8-6 3"))),
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1z"))),
	)
	return x.Svg(svgArgs...)
}
