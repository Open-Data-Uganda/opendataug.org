<script lang="ts">
	import BrandShowcase from '$lib/components/BrandShowcase.svelte';
	import Container from '$lib/components/Container.svelte';
	import FAQ from '$lib/components/FAQ.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
</script>

<Container>
	<div class="bg-white">
		<Navbar />

		<div class="relative isolate px-6 pt-14 lg:px-8">
			<div
				aria-hidden="true"
				class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
			>
				<div
					style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
					class="aspect-1155/678 bg-linear-to-tr relative left-[calc(50%-11rem)] w-[36.125rem] -translate-x-1/2 rotate-[30deg] from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
				></div>
			</div>
			<div class="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
				<div class="hidden sm:mb-8 sm:flex sm:justify-center">
					<div
						class="relative rounded-full px-3 py-1 text-sm/6 text-gray-600 ring-1 ring-gray-900/10 hover:ring-gray-900/20"
					>
						Built for developers
					</div>
				</div>
				<div class="text-center">
					<h1 class="text-balance text-5xl font-semibold tracking-tight text-gray-900 sm:text-6xl">
						Data about Uganda 🇺🇬
					</h1>
					<p class="mt-8 text-pretty text-lg font-medium text-gray-500 sm:text-xl/8">
						Access comprehensive data about Uganda's districts, villages, and administrative units
						through our simple API.
					</p>
					<div class="mt-10 flex items-center justify-center gap-x-6">
						<a
							target="_blank"
							href="https://docs.opendataug.org"
							class="shadow-xs bg-primary hover:bg-primary/80 rounded-md px-3.5 py-2.5 text-sm font-semibold text-white focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
						>
							Read the docs
						</a>
					</div>
				</div>
			</div>

			<div
				aria-hidden="true"
				class="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]"
			>
				<div
					style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
					class="aspect-1155/678 bg-linear-to-tr relative left-[calc(50%+3rem)] w-[36.125rem] -translate-x-1/2 from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
				></div>
			</div>
		</div>

		<BrandShowcase />
		<FAQ />
		<Footer />
	</div>
</Container>
