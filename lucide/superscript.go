package lucide

import x "github.com/bloxui/blox"

// Superscript creates a Superscript Lucide icon.
func Superscript(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-superscript", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m4 19 8-8"))),
		x.Child(x.Path(x.D("m12 19-8-8"))),
		x.Child(x.Path(x.D("M20 12h-4c0-1.5.442-2 1.5-2.5S20 8.334 20 7.002c0-.472-.17-.93-.484-1.29a2.105 2.105 0 0 0-2.617-.436c-.42.239-.738.614-.899 1.06"))),
	)
	return x.Svg(svgArgs...)
}
