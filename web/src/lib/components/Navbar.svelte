<script lang="ts">
	import logo from '$lib/assets/logo.png';

	let mobileMenuOpen = false;

	const navigation = [{ name: 'Docs', href: 'https://docs.opendataug.org' }];
</script>

<header class="inset-x-0 top-0 z-50 mx-auto">
	<nav aria-label="Global" class="flex items-center justify-between">
		<div class="flex lg:flex-1">
			<a href="/" class="">
				<span class="sr-only">Open Data Uganda</span>
				<img alt="" src={logo} class="h-8 w-auto" />
			</a>
		</div>
		<div class="flex lg:hidden">
			<button
				type="button"
				on:click={() => (mobileMenuOpen = true)}
				class="-m-2.5 inline-flex items-center justify-center rounded-md text-gray-700"
			>
				<span class="sr-only">Open main menu</span>
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
				href="https://app.opendataug.org"
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
						<span class="sr-only">Open Data Uganda</span>
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
								href="https://app.opendataug.org"
								class="-mx-3 block rounded-lg px-3 py-2.5 text-base/7 font-semibold text-gray-900 hover:bg-gray-50"
							>
								Sign Up
							</a>
						</div>
						<div class="py-6">
							<a
								href="https://app.opendataug.org/login"
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
