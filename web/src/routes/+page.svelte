<script lang="ts">
	import logo from '$lib/assets/logo.png';
	import BrandShowcase from '$lib/components/BrandShowcase.svelte';
	import Container from '$lib/components/Container.svelte';
	import FAQ from '$lib/components/FAQ.svelte';
	import Footer from '$lib/components/Footer.svelte';

	// Mobile menu state
	let mobileMenuOpen = false;

	const navigation = [{ name: 'Docs', href: 'https://docs.opendataug.org' }];
</script>

<Container>
	<div class="bg-white">
		<header class="inset-x-0 top-0 z-50 mx-auto">
			<nav aria-label="Global" class="flex items-center justify-between p-6 lg:px-8">
				<div class="flex lg:flex-1">
					<a href="/" class="-m-1.5 p-1.5">
						<span class="sr-only">Your Company</span>
						<img alt="" src={logo} class="h-8 w-auto" />
					</a>
				</div>
				<div class="flex lg:hidden">
					<button
						type="button"
						on:click={() => (mobileMenuOpen = true)}
						class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
					>
						<span class="sr-only">Open main menu</span>
						<!-- Replace with Svelte icon -->
						<svg class="size-6" viewBox="0 0 24 24" fill="none" stroke="currentColor">
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M4 6h16M4 12h16M4 18h16"
							/>
						</svg>
					</button>
				</div>
				<div class="hidden lg:flex lg:gap-x-12">
					{#each navigation as item}
						<a href={item.href} class="text-sm/6 font-semibold text-gray-900">
							{item.name}
						</a>
					{/each}
				</div>
				<div class="hidden items-center gap-4 lg:flex lg:flex-1 lg:justify-end">
					<a href="https://app.opendataug.org/login" class="text-sm/6 font-semibold text-gray-900">
						Log in <span aria-hidden="true">→</span>
					</a>

					<a
						href="https://app.opendataug.org/signup"
						class="bg-primary hover:bg-primary/90 rounded-lg px-4 py-2 text-sm/6 font-semibold text-white transition-colors"
					>
						Get Started
					</a>
				</div>
			</nav>

			{#if mobileMenuOpen}
				<div class="lg:hidden">
					<div
						class="fixed inset-0 z-50"
						on:click={() => (mobileMenuOpen = false)}
						on:keydown={(e) => e.key === 'Escape' && (mobileMenuOpen = false)}
						role="button"
						tabindex="0"
					></div>
					<div
						class="fixed inset-y-0 right-0 z-50 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10"
					>
						<div class="flex items-center justify-between">
							<a href="/" class="-m-1.5 p-1.5">
								<span class="sr-only">Your Company</span>
								<img
									alt=""
									src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=600"
									class="h-8 w-auto"
								/>
							</a>
							<button
								type="button"
								on:click={() => (mobileMenuOpen = false)}
								class="-m-2.5 rounded-md p-2.5 text-gray-700"
							>
								<span class="sr-only">Close menu</span>
								<!-- Replace with Svelte icon -->
								<svg class="size-6" viewBox="0 0 24 24" fill="none" stroke="currentColor">
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M6 18L18 6M6 6l12 12"
									/>
								</svg>
							</button>
						</div>
						<div class="mt-6 flow-root">
							<div class="-my-6 divide-y divide-gray-500/10">
								<div class="space-y-2 py-6">
									{#each navigation as item}
										<a
											href={item.href}
											class="-mx-3 block rounded-lg px-3 py-2 text-base/7 font-semibold text-gray-900 hover:bg-gray-50"
										>
											{item.name}
										</a>
									{/each}
								</div>
								<div class="py-6">
									<a
										href="/login"
										class="-mx-3 block rounded-lg px-3 py-2.5 text-base/7 font-semibold text-gray-900 hover:bg-gray-50"
									>
										Log in
									</a>
								</div>
							</div>
						</div>
					</div>
				</div>
			{/if}
		</header>

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
