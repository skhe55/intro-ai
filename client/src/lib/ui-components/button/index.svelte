<script lang="ts">
    export let disabled : boolean = false;
    export let type: "submit" | "reset" | "button" | null = null;
    export let className: string = "";
</script>

<button 
    class={className ? className : "button"}
    type={type} 
    disabled={disabled} 
    on:click
>
    <slot />
</button>