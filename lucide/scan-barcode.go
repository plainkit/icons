package lucide

import x "github.com/bloxui/blox"

// ScanBarcode creates a Scan Barcode Lucide icon.
func ScanBarcode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scan-barcode", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M17 3h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M7 21H5a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Path(x.D("M8 7v10"))),
		x.Child(x.Path(x.D("M12 7v10"))),
		x.Child(x.Path(x.D("M17 7v10"))),
	)
	return x.Svg(svgArgs...)
}
