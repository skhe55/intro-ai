<script lang="ts">
    import Portal from '$lib/portal/index.svelte';
    import { Button } from '$lib/ui-components/index';
	import type { TModalButton } from '$types';

    export let header: string
    export let buttons: TModalButton[]
</script>

<Portal>
    <div class="modal">
        <div class="modal-wrapper">
            <div class="modal-container">
                <header>
                    <h2>{header}</h2>
                </header>
                <div class="modal-container__body">
                    <slot />
                </div>
                <footer>
                    <div class="modal__buttons">
                        {#each buttons as button (button.id)}
                            <Button on:click={button.onClick}>{button.text}</Button>
                        {/each}
                    </div>
                </footer>
            </div>
        </div>
    </div>
</Portal>
