package lucide

import x "github.com/bloxui/blox"

// HandFist creates a Hand Fist Lucide icon.
func HandFist(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hand-fist", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.035 17.012a3 3 0 0 0-3-3l-.311-.002a.72.72 0 0 1-.505-1.229l1.195-1.195A2 2 0 0 1 10.828 11H12a2 2 0 0 0 0-4H9.243a3 3 0 0 0-2.122.879l-2.707 2.707A4.83 4.83 0 0 0 3 14a8 8 0 0 0 8 8h2a8 8 0 0 0 8-8V7a2 2 0 1 0-4 0v2a2 2 0 1 0 4 0"))),
		x.Child(x.Path(x.D("M13.888 9.662A2 2 0 0 0 17 8V5A2 2 0 1 0 13 5"))),
		x.Child(x.Path(x.D("M9 5A2 2 0 1 0 5 5V10"))),
		x.Child(x.Path(x.D("M9 7V4A2 2 0 1 1 13 4V7.268"))),
	)
	return x.Svg(svgArgs...)
}
