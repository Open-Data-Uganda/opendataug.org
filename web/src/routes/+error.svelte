<script>
	import Container from '$lib/components/Container.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
</script>

<Container>
	<Navbar />

	<div
		class="flex h-screen min-h-screen place-content-center items-center justify-center bg-white px-4"
	>
		<h1 class="uppercase tracking-widest text-gray-500">404 | Page Not Found</h1>
	</div>

	<Footer />
</Container>
