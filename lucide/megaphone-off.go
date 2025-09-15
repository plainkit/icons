package lucide

import x "github.com/plainkit/blox"

// MegaphoneOff creates a Megaphone Off Lucide icon.
func MegaphoneOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-megaphone-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11.636 6A13 13 0 0 0 19.4 3.2 1 1 0 0 1 21 4v11.344"))),
		x.Child(x.Path(x.D("M14.378 14.357A13 13 0 0 0 11 14H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M6 14a12 12 0 0 0 2.4 7.2 2 2 0 0 0 3.2-2.4A8 8 0 0 1 10 14"))),
		x.Child(x.Path(x.D("M8 8v6"))),
	)
	return x.Svg(svgArgs...)
}
