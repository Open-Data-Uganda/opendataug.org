<script lang="ts">
	import FeatureCard from './FeatureCard.svelte';
</script>

<section class="mx-auto max-w-7xl px-4 py-16">
	<div class="grid grid-cols-1 items-center gap-16">
		<div class="flex flex-col gap-8">
			<span class="w-fit rounded-full bg-purple-100 px-4 py-2 text-sm font-medium text-purple-600">
				UGANDA DATA API
			</span>

			<h1 class="text-4xl font-bold tracking-tight text-gray-900 lg:text-5xl">
				Access Uganda's Geographic Data with Ease
			</h1>

			<div class="flex flex-col gap-8 md:flex-row">
				<FeatureCard
					icon="verify"
					title="Comprehensive Coverage"
					description="Get detailed data on all districts, villages, and administrative units across Uganda."
				/>

				<FeatureCard
					icon="brand"
					title="Simple Integration"
					description="Easy-to-use API endpoints for seamless access to Uganda's geographic information."
				/>

				<FeatureCard
					icon="boost"
					title="Always Up-to-Date"
					description="Stay current with the latest administrative boundaries and regional data updates."
				/>
			</div>
		</div>
	</div>
</section>

<style>
	h1 {
		font-size: 3rem;
		font-weight: 700;
		line-height: 1.2;
		color: #111827;
	}
</style>
