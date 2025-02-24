<script lang="ts">
	import CTA from './CTA.svelte';

	interface Link {
		label: string;
		href: string;
	}

	const footerLinks: Link[] = [
		{ label: 'About us', href: '/about' },
		{ label: 'Login', href: 'https://app.opendataug.org/login' },
		{ label: 'Get started', href: 'https://app.opendataug.org' },
		{ label: 'Docs ', href: 'https://docs.opendataug.org' },
		{ label: 'Terms', href: 'https://docs.opendataug.org/terms' },
		{ label: 'Privacy Policy', href: 'https://docs.opendataug.org/privacy' }
	];

	const currentYear = new Date().getFullYear();
</script>

<CTA />

<footer class="border-base-content/10 border-b py-20">
	<div class="mx-auto">
		<div class="flex flex-col gap-8">
			<div class="flex flex-col gap-4 lg:mb-20">
				<a href="/" class="flex items-center">
					<span class="text-2xl font-bold">Open Data Uganda</span>
				</a>
				<div class="flex items-center gap-2">
					<a href="mailto:hi@opendataug.org" class="text-base-content/70 hover:text-base-content">
						hi@opendataug.org
					</a>
				</div>
			</div>

			<div class="flex flex-row justify-between gap-6">
				{#each footerLinks as link}
					<a href={link.href} class="text-base-content/70 transition-colors hover:underline">
						{link.label}
					</a>
				{/each}
			</div>
		</div>
	</div>
</footer>

<div class="border-gray-900/10 py-10">
	<p class="text-start text-sm text-gray-900">
		©{currentYear}, Open Data Uganda. All rights reserved
	</p>
</div>
