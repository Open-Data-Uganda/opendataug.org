<script lang="ts">
    export let icon: 'verify' | 'brand' | 'boost' = 'verify';
    export let title: string;
    export let description: string;
</script>

<div class="flex flex-col gap-4 p-4">
    {#if icon === 'verify'}
        <div class="flex h-10 w-10 items-center justify-center rounded-full bg-purple-100 text-purple-600">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
        </div>
    {:else if icon === 'brand'}
        <div class="flex h-10 w-10 items-center justify-center rounded-full bg-blue-100 text-blue-600">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M4 6h16M4 12h16m-7 6h7" />
            </svg>
        </div>
    {:else}
        <div class="flex h-10 w-10 items-center justify-center rounded-full bg-pink-100 text-pink-600">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
            </svg>
        </div>
    {/if}
    
    <h3 class="text-xl font-semibold text-gray-900">{title}</h3>
    <p class="text-gray-500 leading-relaxed">{description}</p>
</div> 