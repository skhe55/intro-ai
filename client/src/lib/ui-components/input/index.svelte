<script lang="ts">
    export let value: string | number;
    export let className: string = "";
</script>

<input class={className ? `input ${className}` : `input`} bind:value={value} />