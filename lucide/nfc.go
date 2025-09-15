package lucide

import x "github.com/plainkit/blox"

// Nfc creates a Nfc Lucide icon.
func Nfc(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-nfc", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 8.32a7.43 7.43 0 0 1 0 7.36"))),
		x.Child(x.Path(x.D("M9.46 6.21a11.76 11.76 0 0 1 0 11.58"))),
		x.Child(x.Path(x.D("M12.91 4.1a15.91 15.91 0 0 1 .01 15.8"))),
		x.Child(x.Path(x.D("M16.37 2a20.16 20.16 0 0 1 0 20"))),
	)
	return x.Svg(svgArgs...)
}
