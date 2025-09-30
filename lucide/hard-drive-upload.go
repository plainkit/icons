package lucide

import (
	html "github.com/plainkit/html"
)

// HardDriveUpload creates a Hard Drive Upload Lucide icon.
func HardDriveUpload(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hard-drive-upload", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 6-4-4-4 4")),
		html.SvgPath(html.AD("M12 2v8")),
		html.SvgRect(html.AWidth("20"), html.AHeight("8"), html.AX("2"), html.AY("14"), html.ARx("2")),
		html.SvgPath(html.AD("M6 18h.01")),
		html.SvgPath(html.AD("M10 18h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
