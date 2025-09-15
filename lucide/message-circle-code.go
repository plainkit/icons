package lucide

import x "github.com/plainkit/blox"

// MessageCircleCode creates a Message Circle Code Lucide icon.
func MessageCircleCode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-message-circle-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 9-3 3 3 3"))),
		x.Child(x.Path(x.D("m14 15 3-3-3-3"))),
		x.Child(x.Path(x.D("M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"))),
	)
	return x.Svg(svgArgs...)
}
