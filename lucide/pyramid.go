package lucide

import x "github.com/bloxui/blox"

// Pyramid creates a Pyramid Lucide icon.
func Pyramid(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-pyramid", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.5 16.88a1 1 0 0 1-.32-1.43l9-13.02a1 1 0 0 1 1.64 0l9 13.01a1 1 0 0 1-.32 1.44l-8.51 4.86a2 2 0 0 1-1.98 0Z"))),
		x.Child(x.Path(x.D("M12 2v20"))),
	)
	return x.Svg(svgArgs...)
}
