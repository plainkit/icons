package lucide

import x "github.com/bloxui/blox"

// BrickWallFire creates a Brick Wall Fire Lucide icon.
func BrickWallFire(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-brick-wall-fire", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 3v2.107"))),
		x.Child(x.Path(x.D("M17 9c1 3 2.5 3.5 3.5 4.5A5 5 0 0 1 22 17a5 5 0 0 1-10 0c0-.3 0-.6.1-.9a2 2 0 1 0 3.3-2C13 11.5 16 9 17 9"))),
		x.Child(x.Path(x.D("M21 8.274V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h3.938"))),
		x.Child(x.Path(x.D("M3 15h5.253"))),
		x.Child(x.Path(x.D("M3 9h8.228"))),
		x.Child(x.Path(x.D("M8 15v6"))),
		x.Child(x.Path(x.D("M8 3v6"))),
	)
	return x.Svg(svgArgs...)
}
