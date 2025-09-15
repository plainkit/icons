package lucide

import x "github.com/plainkit/html"

// DoorClosed creates a Door Closed Lucide icon.
func DoorClosed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-door-closed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 12h.01"))),
		x.Child(x.Path(x.D("M18 20V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14"))),
		x.Child(x.Path(x.D("M2 20h20"))),
	)
	return x.Svg(svgArgs...)
}
