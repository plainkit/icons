package lucide

import x "github.com/bloxui/blox"

// SmartphoneNfc creates a Smartphone Nfc Lucide icon.
func SmartphoneNfc(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-smartphone-nfc", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("1"))),
		x.Child(x.Path(x.D("M13 8.32a7.43 7.43 0 0 1 0 7.36"))),
		x.Child(x.Path(x.D("M16.46 6.21a11.76 11.76 0 0 1 0 11.58"))),
		x.Child(x.Path(x.D("M19.91 4.1a15.91 15.91 0 0 1 .01 15.8"))),
	)
	return x.Svg(svgArgs...)
}
