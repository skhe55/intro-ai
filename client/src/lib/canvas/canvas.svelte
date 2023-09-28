<script lang="ts">
	import { onDestroy, onMount, setContext } from "svelte";
	import type { TDrawFn } from "$lib/types";

    export let width: number = 300;
    export let height: number = 300;
    export let clearFrames: boolean = true;

    let canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D | null;
    let fnsToDraw: TDrawFn[] = [];

    let frameId: number;

    setContext("canvas", {
        addDrawFn: (fn: TDrawFn) => {
            fnsToDraw = [...fnsToDraw, fn]
        },
        removeDrawFn: (fn: TDrawFn) => {
            let index = fnsToDraw.indexOf(fn);
            if(index > -1) {
                fnsToDraw = [...fnsToDraw.splice(index, 1)];
            }
        },
    })

    const draw = (ctx: CanvasRenderingContext2D) => {
        if(clearFrames) {
            ctx.clearRect(0, 0, canvas.width, canvas.height);
        }
        fnsToDraw.forEach(draw => draw(ctx));
        frameId = requestAnimationFrame(() => draw(ctx));
    };

    onMount(() => {
        ctx = canvas.getContext("2d");
        if(ctx) {
            frameId = requestAnimationFrame(() => draw(ctx as CanvasRenderingContext2D))
        }
    });

    onDestroy(() => {
        if(frameId) {
            cancelAnimationFrame(frameId);
        }
    });

    $: {
        console.log(fnsToDraw)
    }
</script>

<canvas 
    bind:this={canvas} {width} {height}
    on:mousemove
    on:mouseleave
    on:mousedown
    on:click
    on:contextmenu={(e) => e.preventDefault()}
></canvas>
<slot />

<style>

</style>