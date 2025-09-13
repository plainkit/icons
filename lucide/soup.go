package lucide

import x "github.com/bloxui/blox"

// Soup creates a Soup Lucide icon.
func Soup(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-soup", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 21a9 9 0 0 0 9-9H3a9 9 0 0 0 9 9Z"))),
		x.Child(x.Path(x.D("M7 21h10"))),
		x.Child(x.Path(x.D("M19.5 12 22 6"))),
		x.Child(x.Path(x.D("M16.25 3c.27.1.8.53.75 1.36-.06.83-.93 1.2-1 2.02-.05.78.34 1.24.73 1.62"))),
		x.Child(x.Path(x.D("M11.25 3c.27.1.8.53.74 1.36-.05.83-.93 1.2-.98 2.02-.06.78.33 1.24.72 1.62"))),
		x.Child(x.Path(x.D("M6.25 3c.27.1.8.53.75 1.36-.06.83-.93 1.2-1 2.02-.05.78.34 1.24.74 1.62"))),
	)
	return x.Svg(svgArgs...)
}
