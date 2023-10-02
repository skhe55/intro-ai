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


<style lang="scss">
    @use '../../../styles/lib/mixins.scss' as *;
	@use '../../../styles/lib/variables.scss' as *;

    .button {
        cursor: pointer;
        
        border-radius: 6px;

        padding: 8px 12px;

        @include text-middle($black, bold);

        transition: all 1s ease-in;
        
        box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;

        border: none;

        &:hover {
            background-color: $link-hover-background-color;
        }
    }
</style>