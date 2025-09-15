package lucide

import x "github.com/plainkit/blox"

// Box creates a Box Lucide icon.
func Box(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-box", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"))),
		x.Child(x.Path(x.D("m3.3 7 8.7 5 8.7-5"))),
		x.Child(x.Path(x.D("M12 22V12"))),
	)
	return x.Svg(svgArgs...)
}
