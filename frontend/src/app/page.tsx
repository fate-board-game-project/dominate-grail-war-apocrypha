import { Button } from "@/components/ui/button";

export default async function Home() {
	const _pong = await fetch(`${process.env.API_SERVER_ENDPOINT}/ping`);
	// console.log(pong);

	return (
		<main>
			<h1>Dominate Grail War - Apocrypha</h1>

			<Button>Login with Google</Button>
		</main>
	);
}
