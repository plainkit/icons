package lucide

import x "github.com/plainkit/html"

// SmartphoneCharging creates a Smartphone Charging Lucide icon.
func SmartphoneCharging(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-smartphone-charging", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("20"), x.X("5"), x.Y("2"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M12.667 8 10 12h4l-2.667 4"))),
	)
	return x.Svg(svgArgs...)
}
