package lucide

import x "github.com/bloxui/blox"

// SignpostBig creates a Signpost Big Lucide icon.
func SignpostBig(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-signpost-big", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 9H4L2 7l2-2h6"))),
		x.Child(x.Path(x.D("M14 5h6l2 2-2 2h-6"))),
		x.Child(x.Path(x.D("M10 22V4a2 2 0 1 1 4 0v18"))),
		x.Child(x.Path(x.D("M8 22h8"))),
	)
	return x.Svg(svgArgs...)
}
