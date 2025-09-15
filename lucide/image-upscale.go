package lucide

import x "github.com/plainkit/html"

// ImageUpscale creates a Image Upscale Lucide icon.
func ImageUpscale(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-image-upscale", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 3h5v5"))),
		x.Child(x.Path(x.D("M17 21h2a2 2 0 0 0 2-2"))),
		x.Child(x.Path(x.D("M21 12v3"))),
		x.Child(x.Path(x.D("m21 3-5 5"))),
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2"))),
		x.Child(x.Path(x.D("m5 21 4.144-4.144a1.21 1.21 0 0 1 1.712 0L13 19"))),
		x.Child(x.Path(x.D("M9 3h3"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("10"), x.X("3"), x.Y("11"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
