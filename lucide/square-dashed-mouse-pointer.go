package lucide

import x "github.com/bloxui/blox"

// SquareDashedMousePointer creates a Square Dashed Mouse Pointer Lucide icon.
func SquareDashedMousePointer(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-dashed-mouse-pointer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.034 12.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.943l-3.444 1.068a1 1 0 0 0-.66.66l-1.067 3.443a.5.5 0 0 1-.943.033z"))),
		x.Child(x.Path(x.D("M5 3a2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M19 3a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M5 21a2 2 0 0 1-2-2"))),
		x.Child(x.Path(x.D("M9 3h1"))),
		x.Child(x.Path(x.D("M9 21h2"))),
		x.Child(x.Path(x.D("M14 3h1"))),
		x.Child(x.Path(x.D("M3 9v1"))),
		x.Child(x.Path(x.D("M21 9v2"))),
		x.Child(x.Path(x.D("M3 14v1"))),
	)
	return x.Svg(svgArgs...)
}
