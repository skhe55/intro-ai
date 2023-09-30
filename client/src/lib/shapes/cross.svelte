<script lang="ts">
	import type { LineConfig } from 'konva/lib/shapes/Line';
	import { Group, Line } from 'svelte-konva';

	export let config: LineConfig;
	let { points } = config;
	let mirroredPoints = points ? [points[2], points[1], points[0], points[3]] : [];
</script>

<Group on:click>
	<Line {config} />
	<Line config={{ ...config, points: mirroredPoints }} />
</Group>
