package lucide

import x "github.com/plainkit/blox"

// Spotlight creates a Spotlight Lucide icon.
func Spotlight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-spotlight", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15.295 19.562 16 22"))),
		x.Child(x.Path(x.D("m17 16 3.758 2.098"))),
		x.Child(x.Path(x.D("m19 12.5 3.026-.598"))),
		x.Child(x.Path(x.D("M7.61 6.3a3 3 0 0 0-3.92 1.3l-1.38 2.79a3 3 0 0 0 1.3 3.91l6.89 3.597a1 1 0 0 0 1.342-.447l3.106-6.211a1 1 0 0 0-.447-1.341z"))),
		x.Child(x.Path(x.D("M8 9V2"))),
	)
	return x.Svg(svgArgs...)
}
