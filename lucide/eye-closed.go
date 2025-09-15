package lucide

import x "github.com/plainkit/blox"

// EyeClosed creates a Eye Closed Lucide icon.
func EyeClosed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-eye-closed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 18-.722-3.25"))),
		x.Child(x.Path(x.D("M2 8a10.645 10.645 0 0 0 20 0"))),
		x.Child(x.Path(x.D("m20 15-1.726-2.05"))),
		x.Child(x.Path(x.D("m4 15 1.726-2.05"))),
		x.Child(x.Path(x.D("m9 18 .722-3.25"))),
	)
	return x.Svg(svgArgs...)
}
