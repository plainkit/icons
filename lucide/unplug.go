package lucide

import x "github.com/plainkit/blox"

// Unplug creates a Unplug Lucide icon.
func Unplug(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-unplug", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m19 5 3-3"))),
		x.Child(x.Path(x.D("m2 22 3-3"))),
		x.Child(x.Path(x.D("M6.3 20.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6-2.3 2.3a2.4 2.4 0 0 0 0 3.4Z"))),
		x.Child(x.Path(x.D("M7.5 13.5 10 11"))),
		x.Child(x.Path(x.D("M10.5 16.5 13 14"))),
		x.Child(x.Path(x.D("m12 6 6 6 2.3-2.3a2.4 2.4 0 0 0 0-3.4l-2.6-2.6a2.4 2.4 0 0 0-3.4 0Z"))),
	)
	return x.Svg(svgArgs...)
}
