package lucide

import x "github.com/bloxui/blox"

// SwatchBook creates a Swatch Book Lucide icon.
func SwatchBook(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-swatch-book", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 17a4 4 0 0 1-8 0V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2Z"))),
		x.Child(x.Path(x.D("M16.7 13H19a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H7"))),
		x.Child(x.Path(x.D("M 7 17h.01"))),
		x.Child(x.Path(x.D("m11 8 2.3-2.3a2.4 2.4 0 0 1 3.404.004L18.6 7.6a2.4 2.4 0 0 1 .026 3.434L9.9 19.8"))),
	)
	return x.Svg(svgArgs...)
}
