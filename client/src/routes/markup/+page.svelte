<script lang="ts">
	import Canvas from "$lib/canvas/canvas.svelte";
    import Square from "$lib/shapes/square.svelte";
	import type { TSquare } from "$lib/types";
	import { afterUpdate } from "svelte";

    let uuid = crypto.randomUUID();

    let squares: TSquare[] = [];
    let squareCoords: Omit<TSquare, "id"> = {
        x_top: [0, 0],
        x_bottom: [0, 0],
        y_top: [0, 0],
        y_bottom: [0, 0],
    };

    let isMarkupMode: boolean = false;

    const followMouse = (e: MouseEvent) => {
        if(isMarkupMode) {           
            squareCoords = {
                x_top: squareCoords.x_top,
                x_bottom: [squareCoords.x_top[0], e.offsetY],
                y_top: [e.offsetX, squareCoords.x_top[1]],
                y_bottom: [e.offsetX, e.offsetY],
            }
        }
    };

    const onStartMarking = (e: MouseEvent) => {
        if(isMarkupMode) {
            isMarkupMode = false;
            squares = [...squares, {...squareCoords, id: uuid}];
        } else {
            squareCoords = {
                x_top: [e.offsetX, e.offsetY],
                x_bottom: [e.offsetX, e.offsetY],
                y_top: [e.offsetX, e.offsetY],
                y_bottom: [e.offsetX, e.offsetY],
            };
            isMarkupMode = true;
        }
    };

    const onCancelMarkupMode = (e: MouseEvent) => {
        if(e.button === 2) {
            isMarkupMode = false;
        } 
    };
    
    afterUpdate(() => {
        uuid = crypto.randomUUID();
    });
    $: console.log(squares)
</script>
<section class="markup-page">
    <header>
        <h1>Markup page</h1>
    </header>
    <main>
        <div class="canvas-container">
            <Canvas 
                width={600}
                height={600}
                on:mousemove={followMouse}
                on:mousedown={onCancelMarkupMode}
                on:click={onStartMarking}
            >
                {#each squares as square}
                    <Square 
                        x_top={square.x_top}
                        x_bottom={square.x_bottom}
                        y_top={square.y_top}
                        y_bottom={square.y_bottom}
                    />
                {/each}
                <Square 
                    x_top={squareCoords.x_top} 
                    x_bottom={squareCoords.x_bottom}
                    y_top={squareCoords.y_top} 
                    y_bottom={squareCoords.y_bottom} 
                />
            </Canvas>
        </div>
    </main>
</section>

<style lang="scss">
    .markup-page {
        display: flex;
        flex-direction: column;

        align-items: center;

        gap: 20px;
        margin: 20px;

        height: 100%;

        main {
            .canvas-container {
                box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
            }
        }
    }
</style>