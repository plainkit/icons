package lucide

import x "github.com/bloxui/blox"

// ScanQrCode creates a Scan Qr Code Lucide icon.
func ScanQrCode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scan-qr-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 12v4a1 1 0 0 1-1 1h-4"))),
		x.Child(x.Path(x.D("M17 3h2a2 2 0 0 1 2 2v2"))),
		x.Child(x.Path(x.D("M17 8V7"))),
		x.Child(x.Path(x.D("M21 17v2a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M3 7V5a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M7 17h.01"))),
		x.Child(x.Path(x.D("M7 21H5a2 2 0 0 1-2-2v-2"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("5"), x.X("7"), x.Y("7"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
