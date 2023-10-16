<script lang="ts">
    export let value: string | number;
    export let className: string = "";
</script>

<input class={className ? className : "input"} bind:value={value} />

<style lang="scss">
    @use '../../../styles/lib/variables.scss' as *;

    .input {
        border: 1px solid $link-hover-background-color;
        border-radius: 8px;
        outline: none;
        padding: 8px 12px;
    }
</style>