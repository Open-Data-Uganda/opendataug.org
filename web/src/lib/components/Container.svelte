<script lang="ts">
	export let className = '';
</script>

<div class="mx-auto w-full max-w-7xl px-4 lg:px-0">
	<div class="py-8 {className}">
		<slot />
	</div>
</div>
