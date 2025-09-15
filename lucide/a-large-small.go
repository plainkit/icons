package lucide

import x "github.com/plainkit/blox"

// ALargeSmall creates a A Large Small Lucide icon.
func ALargeSmall(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-a-large-small", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 16 2.536-7.328a1.02 1.02 1 0 1 1.928 0L22 16"))),
		x.Child(x.Path(x.D("M15.697 14h5.606"))),
		x.Child(x.Path(x.D("m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"))),
		x.Child(x.Path(x.D("M3.304 13h6.392"))),
	)
	return x.Svg(svgArgs...)
}
