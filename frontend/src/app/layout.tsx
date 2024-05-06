import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
	title: "Dominate Grail War - Apocrypha",
	description: "Dominate Grail War - Apocrypha",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="ja">
			<body>{children}</body>
		</html>
	);
}
