package lucide

import x "github.com/bloxui/blox"

// BrickWallShield creates a Brick Wall Shield Lucide icon.
func BrickWallShield(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-brick-wall-shield", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 9v1.258"))),
		x.Child(x.Path(x.D("M16 3v5.46"))),
		x.Child(x.Path(x.D("M21 9.118V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h5.75"))),
		x.Child(x.Path(x.D("M22 17.5c0 2.499-1.75 3.749-3.83 4.474a.5.5 0 0 1-.335-.005c-2.085-.72-3.835-1.97-3.835-4.47V14a.5.5 0 0 1 .5-.499c1 0 2.25-.6 3.12-1.36a.6.6 0 0 1 .76-.001c.875.765 2.12 1.36 3.12 1.36a.5.5 0 0 1 .5.5z"))),
		x.Child(x.Path(x.D("M3 15h7"))),
		x.Child(x.Path(x.D("M3 9h12.142"))),
		x.Child(x.Path(x.D("M8 15v6"))),
		x.Child(x.Path(x.D("M8 3v6"))),
	)
	return x.Svg(svgArgs...)
}
