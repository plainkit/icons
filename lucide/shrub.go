package lucide

import x "github.com/plainkit/blox"

// Shrub creates a Shrub Lucide icon.
func Shrub(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shrub", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 22v-5.172a2 2 0 0 0-.586-1.414L9.5 13.5"))),
		x.Child(x.Path(x.D("M14.5 14.5 12 17"))),
		x.Child(x.Path(x.D("M17 8.8A6 6 0 0 1 13.8 20H10A6.5 6.5 0 0 1 7 8a5 5 0 0 1 10 0z"))),
	)
	return x.Svg(svgArgs...)
}
