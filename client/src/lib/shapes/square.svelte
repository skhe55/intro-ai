<script lang="ts">
	import { getContext, onDestroy, onMount } from "svelte";
	import type { TCanvasContext, TPoint } from "../types";


    export let x_top: TPoint;
    export let x_bottom: TPoint;
    export let y_top: TPoint;
    export let y_bottom: TPoint;

    let canvasContext: TCanvasContext = getContext("canvas");

    const draw = (ctx: CanvasRenderingContext2D) => {
        ctx.beginPath();

        ctx.moveTo(...x_top);
        ctx.lineTo(...x_bottom);
        ctx.stroke();

        ctx.moveTo(...x_top);
        ctx.lineTo(...y_top);
        ctx.stroke();

        ctx.moveTo(...y_top);
        ctx.lineTo(...y_bottom);
        ctx.stroke();

        ctx.moveTo(...x_bottom);
        ctx.lineTo(...y_bottom);
        ctx.stroke();
    };

    onMount(() => {
        canvasContext.addDrawFn(draw);
    });

    onDestroy(() => {
        canvasContext.removeDrawFn(draw);
    });
</script>