package lucide

import x "github.com/bloxui/blox"

// DoorClosedLocked creates a Door Closed Locked Lucide icon.
func DoorClosedLocked(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-door-closed-locked", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 12h.01"))),
		x.Child(x.Path(x.D("M18 9V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14"))),
		x.Child(x.Path(x.D("M2 20h8"))),
		x.Child(x.Path(x.D("M20 17v-2a2 2 0 1 0-4 0v2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("5"), x.X("14"), x.Y("17"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
