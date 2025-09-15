package lucide

import x "github.com/plainkit/html"

// ReceiptJapaneseYen creates a Receipt Japanese Yen Lucide icon.
func ReceiptJapaneseYen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-receipt-japanese-yen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		x.Child(x.Path(x.D("m12 10 3-3"))),
		x.Child(x.Path(x.D("m9 7 3 3v7.5"))),
		x.Child(x.Path(x.D("M9 11h6"))),
		x.Child(x.Path(x.D("M9 15h6"))),
	)
	return x.Svg(svgArgs...)
}
