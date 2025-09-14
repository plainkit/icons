package lucide

import x "github.com/bloxui/blox"

// Barcode creates a Barcode Lucide icon.
func Barcode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-barcode", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 5v14"))),
		x.Child(x.Path(x.D("M8 5v14"))),
		x.Child(x.Path(x.D("M12 5v14"))),
		x.Child(x.Path(x.D("M17 5v14"))),
		x.Child(x.Path(x.D("M21 5v14"))),
	)
	return x.Svg(svgArgs...)
}
