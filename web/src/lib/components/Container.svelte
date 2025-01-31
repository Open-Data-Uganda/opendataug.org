<script lang="ts">
    export let className = "";
</script>

<div class="mx-auto w-full max-w-7xl px-4 sm:px-6 lg:px-8">
    <div class="py-8 {className}">
        <slot />
    </div>
</div> 