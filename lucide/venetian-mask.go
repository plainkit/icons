package lucide

import x "github.com/bloxui/blox"

// VenetianMask creates a Venetian Mask Lucide icon.
func VenetianMask(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-venetian-mask", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 11c-1.5 0-2.5.5-3 2"))),
		x.Child(x.Path(x.D("M4 6a2 2 0 0 0-2 2v4a5 5 0 0 0 5 5 8 8 0 0 1 5 2 8 8 0 0 1 5-2 5 5 0 0 0 5-5V8a2 2 0 0 0-2-2h-3a8 8 0 0 0-5 2 8 8 0 0 0-5-2z"))),
		x.Child(x.Path(x.D("M6 11c1.5 0 2.5.5 3 2"))),
	)
	return x.Svg(svgArgs...)
}
