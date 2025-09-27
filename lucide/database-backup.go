package lucide

import (
	html "github.com/plainkit/html"
)

// DatabaseBackup creates a Database Backup Lucide icon.
func DatabaseBackup(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-database-backup", args)
	children := []html.SvgArg{
		html.Child(html.SvgEllipse(html.ACx("12"), html.ACy("5"), html.ARx("9"), html.ARy("3"))),
		html.Child(html.SvgPath(html.AD("M3 12a9 3 0 0 0 5 2.69"))),
		html.Child(html.SvgPath(html.AD("M21 9.3V5"))),
		html.Child(html.SvgPath(html.AD("M3 5v14a9 3 0 0 0 6.47 2.88"))),
		html.Child(html.SvgPath(html.AD("M12 12v4h4"))),
		html.Child(html.SvgPath(html.AD("M13 20a5 5 0 0 0 9-3 4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L12 16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
