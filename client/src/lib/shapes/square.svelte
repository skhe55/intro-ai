<script lang="ts">
	import { Group, Rect, type KonvaMouseEvent } from 'svelte-konva';
	import Cross from './cross.svelte';
	import type { RectConfig } from 'konva/lib/shapes/Rect';
	import type { LineConfig } from 'konva/lib/shapes/Line';

	export let rectConfig: RectConfig,
		crossConfig: LineConfig,
		crossOnClick: (e: KonvaMouseEvent) => void,
		isShowCross: boolean;
</script>

<Group>
	<Rect config={rectConfig} on:mouseenter on:mouseleave/>
	{#if isShowCross}
		<Cross config={crossConfig} on:click={crossOnClick} />
	{/if}
</Group>