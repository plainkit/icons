package lucide

import x "github.com/plainkit/blox"

// Hamburger creates a Hamburger Lucide icon.
func Hamburger(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hamburger", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 16H4a2 2 0 1 1 0-4h16a2 2 0 1 1 0 4h-4.25"))),
		x.Child(x.Path(x.D("M5 12a2 2 0 0 1-2-2 9 7 0 0 1 18 0 2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M5 16a2 2 0 0 0-2 2 3 3 0 0 0 3 3h12a3 3 0 0 0 3-3 2 2 0 0 0-2-2q0 0 0 0"))),
		x.Child(x.Path(x.D("m6.67 12 6.13 4.6a2 2 0 0 0 2.8-.4l3.15-4.2"))),
	)
	return x.Svg(svgArgs...)
}
