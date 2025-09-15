package lucide

import x "github.com/plainkit/html"

// QrCode creates a Qr Code Lucide icon.
func QrCode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-qr-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("5"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("5"), x.X("16"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("5"), x.X("3"), x.Y("16"), x.Rx("1"))),
		x.Child(x.Path(x.D("M21 16h-3a2 2 0 0 0-2 2v3"))),
		x.Child(x.Path(x.D("M21 21v.01"))),
		x.Child(x.Path(x.D("M12 7v3a2 2 0 0 1-2 2H7"))),
		x.Child(x.Path(x.D("M3 12h.01"))),
		x.Child(x.Path(x.D("M12 3h.01"))),
		x.Child(x.Path(x.D("M12 16v.01"))),
		x.Child(x.Path(x.D("M16 12h1"))),
		x.Child(x.Path(x.D("M21 12v.01"))),
		x.Child(x.Path(x.D("M12 21v-1"))),
	)
	return x.Svg(svgArgs...)
}
