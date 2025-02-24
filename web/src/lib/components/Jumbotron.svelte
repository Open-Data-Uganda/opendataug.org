<script>
	import { ChevronRight, Home } from 'lucide-svelte';

	let { title, currentPage } = $props();
</script>

<section class="relative overflow-hidden bg-gradient-to-br from-indigo-100 to-white">
	<!-- SVG Pattern -->
	<div class="absolute inset-0 -z-10 opacity-[0.15]">
		<svg
			class="h-full w-full"
			xmlns="http://www.w3.org/2000/svg"
			width="100%"
			height="100%"
			fill="none"
			viewBox="0 0 80 80"
		>
			<pattern
				id="pattern-squares"
				x="0"
				y="0"
				width="20"
				height="20"
				patternUnits="userSpaceOnUse"
			>
				<path d="M0 0L20 0L20 20L0 20L0 0Z" fill="none" stroke="currentColor" stroke-width="0.5" />
				<circle cx="10" cy="10" r="1" fill="currentColor" />
			</pattern>
			<rect width="100%" height="100%" fill="url(#pattern-squares)" />
		</svg>
	</div>

	<div class="mx-auto max-w-7xl px-6 py-24 sm:py-32 lg:px-8">
		<!-- Breadcrumbs -->
		<nav class="mb-8" aria-label="breadcrumb">
			<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm">
				<li class="inline-flex items-center gap-1.5">
					<a href="/" class="hover:text-foreground transition-colors">
						<Home class="h-4 w-4" />
						<span class="sr-only">Home</span>
					</a>
				</li>
				<li>
					<ChevronRight class="text-muted-foreground/40 h-4 w-4" />
				</li>
				<li>
					<a href="/products" class="hover:text-foreground transition-colors"> Home </a>
				</li>
				<li>
					<ChevronRight class="text-muted-foreground/40 h-4 w-4" />
				</li>
				<li>
					<span class="text-foreground font-medium">{currentPage}</span>
				</li>
			</ol>
		</nav>

		<div class="mx-auto max-w-2xl text-center">
			<h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
				{title}
			</h1>
			<p class="mt-6 text-lg leading-8 text-gray-600"></p>
		</div>
	</div>

	<div
		class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
		aria-hidden="true"
	>
		<div
			class="from-primary to-primary-foreground relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr opacity-20 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
			style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
		></div>
	</div>
</section>
