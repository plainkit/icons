package lucide

import x "github.com/bloxui/blox"

// PencilOff creates a Pencil Off Lucide icon.
func PencilOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-pencil-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 10-6.157 6.162a2 2 0 0 0-.5.833l-1.322 4.36a.5.5 0 0 0 .622.624l4.358-1.323a2 2 0 0 0 .83-.5L14 13.982"))),
		x.Child(x.Path(x.D("m12.829 7.172 4.359-4.346a1 1 0 1 1 3.986 3.986l-4.353 4.353"))),
		x.Child(x.Path(x.D("m15 5 4 4"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
