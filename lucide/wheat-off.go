package lucide

import x "github.com/plainkit/blox"

// WheatOff creates a Wheat Off Lucide icon.
func WheatOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-wheat-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 22 10-10"))),
		x.Child(x.Path(x.D("m16 8-1.17 1.17"))),
		x.Child(x.Path(x.D("M3.47 12.53 5 11l1.53 1.53a3.5 3.5 0 0 1 0 4.94L5 19l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Z"))),
		x.Child(x.Path(x.D("m8 8-.53.53a3.5 3.5 0 0 0 0 4.94L9 15l1.53-1.53c.55-.55.88-1.25.98-1.97"))),
		x.Child(x.Path(x.D("M10.91 5.26c.15-.26.34-.51.56-.73L13 3l1.53 1.53a3.5 3.5 0 0 1 .28 4.62"))),
		x.Child(x.Path(x.D("M20 2h2v2a4 4 0 0 1-4 4h-2V6a4 4 0 0 1 4-4Z"))),
		x.Child(x.Path(x.D("M11.47 17.47 13 19l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L5 19l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"))),
		x.Child(x.Path(x.D("m16 16-.53.53a3.5 3.5 0 0 1-4.94 0L9 15l1.53-1.53a3.49 3.49 0 0 1 1.97-.98"))),
		x.Child(x.Path(x.D("M18.74 13.09c.26-.15.51-.34.73-.56L21 11l-1.53-1.53a3.5 3.5 0 0 0-4.62-.28"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
