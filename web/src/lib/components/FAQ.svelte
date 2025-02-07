<script lang="ts">
	let expandedIndex: number | null = null;

	const faqItems = [
		{
			question: 'What kind of data can I access through the API?',
			answer:
				"Our API provides comprehensive data about Uganda's administrative units including districts, counties, sub-counties, parishes, and villages. You can access demographic information, geographical coordinates, and administrative boundaries."
		},
		{
			question: 'How up-to-date is the data?',
			answer:
				'We regularly update our database to reflect the latest administrative changes and demographic information. Each data point includes a timestamp of its last update, ensuring you know exactly how current the information is.'
		},
		{
			question: 'Do I need an API key to access the data?',
			answer:
				"Yes, you'll need an API key to access the data. You can get one by signing up for a free account. We offer different tiers based on usage requirements."
		},
		{
			question: 'Is there a rate limit for API calls?',
			answer:
				'Yes, we have rate limits in place to ensure fair usage. Free tier accounts have a limit of 1000 requests per day. For higher limits, please check our premium plans.'
		},
		{
			question: 'What format is the data available in?',
			answer:
				'The data is available in JSON format by default. All responses follow standard REST API conventions.'
		}
	];

	function handleToggle(index: number) {
		expandedIndex = expandedIndex === index ? null : index;
	}
</script>

<div
	class="mx-auto my-8 flex flex-col gap-12 bg-gray-100 px-8 py-24 md:flex-row lg:my-16 lg:rounded-lg"
>
	<div class="flex basis-1/2 flex-col pl-20 text-left">
		<p class="mb-4 inline-block font-semibold">FAQs</p>
		<p class="text-base-content text-3xl font-extrabold sm:text-4xl">Get to know us better.</p>
	</div>

	<ul class="basis-1/2 lg:pr-20">
		{#each faqItems as item, index}
			<li>
				<div class="faq-item">
					<button
						class="border-base-content/10 relative flex w-full items-center gap-2 border-t py-5 text-left text-base font-semibold md:text-lg"
						on:click={() => handleToggle(index)}
						aria-expanded={expandedIndex === index}
					>
						<span class="text-base-content flex-1">{item.question}</span>
						<svg
							class="ml-auto h-4 w-4 flex-shrink-0 fill-current"
							viewBox="0 0 16 16"
							xmlns="http://www.w3.org/2000/svg"
						>
							<rect
								y="7"
								width="16"
								height="2"
								rx="1"
								class="transform transition duration-200 ease-out"
							/>
							<rect
								y="7"
								width="16"
								height="2"
								rx="1"
								class="origin-center transform transition duration-200 ease-out {expandedIndex ===
								index
									? 'rotate-0'
									: 'rotate-90'}"
							/>
						</svg>
					</button>
					<div
						class="overflow-hidden transition-all duration-300 ease-in-out"
						style="max-height: {expandedIndex === index
							? '1000px'
							: '0'}; transition: max-height 0.3s ease-in-out"
					>
						<div class="pb-5 leading-relaxed">
							<div class="space-y-2 leading-relaxed">{item.answer}</div>
						</div>
					</div>
				</div>
			</li>
		{/each}
	</ul>
</div>
