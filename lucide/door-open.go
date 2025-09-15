package lucide

import x "github.com/plainkit/blox"

// DoorOpen creates a Door Open Lucide icon.
func DoorOpen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-door-open", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 20H2"))),
		x.Child(x.Path(x.D("M11 4.562v16.157a1 1 0 0 0 1.242.97L19 20V5.562a2 2 0 0 0-1.515-1.94l-4-1A2 2 0 0 0 11 4.561z"))),
		x.Child(x.Path(x.D("M11 4H8a2 2 0 0 0-2 2v14"))),
		x.Child(x.Path(x.D("M14 12h.01"))),
		x.Child(x.Path(x.D("M22 20h-3"))),
	)
	return x.Svg(svgArgs...)
}
