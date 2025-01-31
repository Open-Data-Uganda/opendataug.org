<script lang="ts">
	import '../app.css';
</script>

<slot />

<style>
	:global(body) {
		margin: 0;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
			Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
	}

	:global(*) {
		box-sizing: border-box;
	}
</style>
