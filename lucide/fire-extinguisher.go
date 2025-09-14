package lucide

import x "github.com/bloxui/blox"

// FireExtinguisher creates a Fire Extinguisher Lucide icon.
func FireExtinguisher(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-fire-extinguisher", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 6.5V3a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3.5"))),
		x.Child(x.Path(x.D("M9 18h8"))),
		x.Child(x.Path(x.D("M18 3h-3"))),
		x.Child(x.Path(x.D("M11 3a6 6 0 0 0-6 6v11"))),
		x.Child(x.Path(x.D("M5 13h4"))),
		x.Child(x.Path(x.D("M17 10a4 4 0 0 0-8 0v10a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2Z"))),
	)
	return x.Svg(svgArgs...)
}
