package lucide

import x "github.com/plainkit/blox"

// Spool creates a Spool Lucide icon.
func Spool(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-spool", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 13.44 4.442 17.082A2 2 0 0 0 4.982 21H19a2 2 0 0 0 .558-3.921l-1.115-.32A2 2 0 0 1 17 14.837V7.66"))),
		x.Child(x.Path(x.D("m7 10.56 12.558-3.642A2 2 0 0 0 19.018 3H5a2 2 0 0 0-.558 3.921l1.115.32A2 2 0 0 1 7 9.163v7.178"))),
	)
	return x.Svg(svgArgs...)
}
