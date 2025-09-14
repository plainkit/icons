package lucide

import x "github.com/bloxui/blox"

// Fence creates a Fence Lucide icon.
func Fence(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-fence", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 3 2 5v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"))),
		x.Child(x.Path(x.D("M6 8h4"))),
		x.Child(x.Path(x.D("M6 18h4"))),
		x.Child(x.Path(x.D("m12 3-2 2v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"))),
		x.Child(x.Path(x.D("M14 8h4"))),
		x.Child(x.Path(x.D("M14 18h4"))),
		x.Child(x.Path(x.D("m20 3-2 2v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"))),
	)
	return x.Svg(svgArgs...)
}
