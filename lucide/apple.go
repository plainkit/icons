package lucide

import x "github.com/bloxui/blox"

// Apple creates a Apple Lucide icon.
func Apple(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-apple", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6.528V3a1 1 0 0 1 1-1h0"))),
		x.Child(x.Path(x.D("M18.237 21A15 15 0 0 0 22 11a6 6 0 0 0-10-4.472A6 6 0 0 0 2 11a15.1 15.1 0 0 0 3.763 10 3 3 0 0 0 3.648.648 5.5 5.5 0 0 1 5.178 0A3 3 0 0 0 18.237 21"))),
	)
	return x.Svg(svgArgs...)
}
