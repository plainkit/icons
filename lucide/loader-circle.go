package lucide

import x "github.com/bloxui/blox"

// LoaderCircle creates a Loader Circle Lucide icon.
func LoaderCircle(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-loader-circle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 12a9 9 0 1 1-6.219-8.56"))),
	)
	return x.Svg(svgArgs...)
}
