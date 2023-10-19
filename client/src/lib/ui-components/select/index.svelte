<script lang="ts">
    export let data: any[];
    export let placeholder: string = "please select item";
    export let textField: string;
    export let valueField: string;
    let selected: any;
</script>

<select class="select" bind:value={selected} autocomplete="off">
    <option value="">{placeholder}</option>
    {#each data as elem (elem.id)}
        <option value={elem[valueField]}>{elem[textField]}</option>
    {/each}
</select>
