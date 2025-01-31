<script lang="ts">
    export let label: string = "Click me";
    export let variant: "primary" | "secondary" = "primary";
    export let disabled: boolean = false;
</script>

<button 
    class={variant} 
    on:click 
    {disabled}
>
    {label}
    <slot />
</button>

<style>
    button {
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 1rem;
        transition: background-color 0.2s;
    }

    button:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    .primary {
        background-color: #4CAF50;
        color: white;
    }

    .primary:hover:not(:disabled) {
        background-color: #45a049;
    }

    .secondary {
        background-color: #f0f0f0;
        color: #333;
        border: 1px solid #ccc;
    }

    .secondary:hover:not(:disabled) {
        background-color: #e0e0e0;
    }
</style> 