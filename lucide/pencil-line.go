package lucide

import x "github.com/bloxui/blox"

// PencilLine creates a Pencil Line Lucide icon.
func PencilLine(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-pencil-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 21h8"))),
		x.Child(x.Path(x.D("m15 5 4 4"))),
		x.Child(x.Path(x.D("M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"))),
	)
	return x.Svg(svgArgs...)
}
