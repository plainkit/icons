package lucide

import x "github.com/bloxui/blox"

// TouchpadOff creates a Touchpad Off Lucide icon.
func TouchpadOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-touchpad-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20v-6"))),
		x.Child(x.Path(x.D("M19.656 14H22"))),
		x.Child(x.Path(x.D("M2 14h12"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2"))),
		x.Child(x.Path(x.D("M9.656 4H20a2 2 0 0 1 2 2v10.344"))),
	)
	return x.Svg(svgArgs...)
}
