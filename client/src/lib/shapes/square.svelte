<script lang="ts">
	import { Group, Rect, type KonvaMouseEvent } from 'svelte-konva';
	import Cross from './cross.svelte';
	import type { RectConfig } from 'konva/lib/shapes/Rect';
	import type { LineConfig } from 'konva/lib/shapes/Line';

	export let rectConfig: RectConfig,
		crossConfig: LineConfig,
		crossOnClick: (e: KonvaMouseEvent) => void;
</script>

<Group>
	<Rect config={rectConfig} />
	<Cross config={crossConfig} on:click={crossOnClick} />
</Group>
