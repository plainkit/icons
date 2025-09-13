package lucide

import x "github.com/bloxui/blox"

// LoaderPinwheel creates a Loader Pinwheel Lucide icon.
func LoaderPinwheel(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-loader-pinwheel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 12a1 1 0 0 1-10 0 1 1 0 0 0-10 0"))),
		x.Child(x.Path(x.D("M7 20.7a1 1 0 1 1 5-8.7 1 1 0 1 0 5-8.6"))),
		x.Child(x.Path(x.D("M7 3.3a1 1 0 1 1 5 8.6 1 1 0 1 0 5 8.6"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
