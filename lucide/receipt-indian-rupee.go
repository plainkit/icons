package lucide

import x "github.com/bloxui/blox"

// ReceiptIndianRupee creates a Receipt Indian Rupee Lucide icon.
func ReceiptIndianRupee(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-receipt-indian-rupee", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("M8 7h8"))),
		x.Child(x.Path(x.D("M12 17.5 8 15h1a4 4 0 0 0 0-8"))),
		x.Child(x.Path(x.D("M8 11h8"))),
	)
	return x.Svg(svgArgs...)
}
