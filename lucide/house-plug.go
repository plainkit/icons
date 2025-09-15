package lucide

import x "github.com/plainkit/html"

// HousePlug creates a House Plug Lucide icon.
func HousePlug(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-house-plug", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 12V8.964"))),
		x.Child(x.Path(x.D("M14 12V8.964"))),
		x.Child(x.Path(x.D("M15 12a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-2a1 1 0 0 1 1-1z"))),
		x.Child(x.Path(x.D("M8.5 21H5a2 2 0 0 1-2-2v-9a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2v-2"))),
	)
	return x.Svg(svgArgs...)
}
