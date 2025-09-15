package lucide

import x "github.com/plainkit/blox"

// KeyboardOff creates a Keyboard Off Lucide icon.
func KeyboardOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-keyboard-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M 20 4 A2 2 0 0 1 22 6"))),
		x.Child(x.Path(x.D("M 22 6 L 22 16.41"))),
		x.Child(x.Path(x.D("M 7 16 L 16 16"))),
		x.Child(x.Path(x.D("M 9.69 4 L 20 4"))),
		x.Child(x.Path(x.D("M14 8h.01"))),
		x.Child(x.Path(x.D("M18 8h.01"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2"))),
		x.Child(x.Path(x.D("M6 8h.01"))),
		x.Child(x.Path(x.D("M8 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
