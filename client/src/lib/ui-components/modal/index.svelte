<script lang="ts">
    import Portal from '$lib/portal/index.svelte';
</script>

<Portal>
    <div class="modal">
        <div class="modal-wrapper">
           <slot /> 
        </div>
    </div>
</Portal>
