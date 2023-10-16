<script lang="ts">
	import { afterUpdate, onMount } from 'svelte';
    import { Stage, Layer, Rect, Image, type KonvaMouseEvent } from 'svelte-konva';

	import type { TSquare } from '$lib/types';
	import { initialSquareCoordinates } from '$constants/index';
	import Square from '$lib/shapes/square.svelte';
	import { getPointsUpperRightCorner } from '$lib/utils';
	import { Button } from '$lib/ui-components';

	let uuid = crypto.randomUUID();

	let squares: TSquare[] = [];
	let squareCoords: Omit<TSquare, 'id'> = initialSquareCoordinates;

	let isMarkupMode: boolean = false;
	let isWatchMode: boolean = false;
	let isEditMode: boolean = false;

	let image: HTMLOrSVGImageElement;

	const onToggleMode = () => {
		isEditMode = !isEditMode;
	};

	const onMouseMove = (e: KonvaMouseEvent) => {
		const { evt } = e.detail;
		if (isMarkupMode) {
			squareCoords = {
				x_top: squareCoords.x_top,
				x_bottom: [squareCoords.x_top[0], evt.offsetY],
				y_top: [evt.offsetX, squareCoords.x_top[1]],
				y_bottom: [evt.offsetX, evt.offsetY]
			};
		}
	};

	const onStartMarking = (e: KonvaMouseEvent) => {
		const { evt } = e.detail;
		if (isMarkupMode) {
			squares = [...squares, { ...squareCoords, id: uuid }];
			isMarkupMode = false;
		} else if (!isMarkupMode && !isWatchMode) {
			squareCoords = {
				x_top: [evt.offsetX, evt.offsetY],
				x_bottom: [evt.offsetX, evt.offsetY],
				y_top: [evt.offsetX, evt.offsetY],
				y_bottom: [evt.offsetX, evt.offsetY]
			};
			isMarkupMode = true;
		}
	};

	const onRemoveItem = (_: KonvaMouseEvent, id: string) => {
		squares = squares.filter((item) => item.id !== id);
	};

	onMount(() => {
		const img = document.createElement('img') as HTMLImageElement;
		img.src =
			'https://roblouie.com/wp-content/uploads/2020/04/60788338_304920937106527_8424495022080625603_n.jpg';
		img.onload = () => {
			image = img;
		};

		window.addEventListener('keydown', (e) => {
			if (e.code === 'Escape') {
				isMarkupMode = false;
			} else if (e.code === 'KeyN') {
				isWatchMode = !isWatchMode;
			}
		});
	});

	afterUpdate(() => {
		uuid = crypto.randomUUID();
	});
	$: console.log(squares);
</script>

<section class="markup-page">
	<div class="markup-view-container">
		<div class="markup-controls">
			<div class="markup-controls__buttons">
				<Button>Save</Button>
				<Button>Create a project</Button>
				<Button on:click={onToggleMode}>{isEditMode ? "Disable edit" : "Enable edit"}</Button>
			</div>
			<h3>Current mode: {isWatchMode ? "Watch" : "Markup"}</h3>
		</div>
		<div class="canvas-container">
			<Stage
				config={{ width: 600, height: 600 }}
				on:mousemove={onMouseMove}
				on:click={onStartMarking}
			>
				<Layer>
					<Image config={{ image: image, width: 600, height: 600 }} />
					{#if isMarkupMode}
						<Rect
							config={{
								x: squareCoords.x_top[0],
								y: squareCoords.x_top[1],
								width: squareCoords.y_top[0] - squareCoords.x_top[0],
								height: squareCoords.y_bottom[1] - squareCoords.x_top[1],
								stroke: 'black',
								strokeWidth: 3
							}}
						/>
					{/if}
					{#each squares as square (square.id)}
						<Square
							rectConfig={{
								x: square.x_top[0],
								y: square.x_top[1],
								width: square.y_top[0] - square.x_top[0],
								height: square.y_bottom[1] - square.x_top[1],
								stroke: 'black',
								strokeWidth: 3
							}}
							crossConfig={{
								points: getPointsUpperRightCorner(square),
								stroke: 'red',
								strokeWidth: 4
							}}
							crossOnClick={(e) => onRemoveItem(e, square.id)}
							isShowCross={isEditMode}
						/>
					{/each}
				</Layer>
			</Stage>
		</div>
		<div class="markup-view-tips">
			<h2>How to use it</h2>
			<p>First, we need upload a image that we will mark up.</p>
			<p>Then left-click on image and start mark up.</p>
			<p>After we finished markup, we need saved our results.</p>
		</div>
	</div>
	<div class="markup-storage">
		
	</div>
</section>

<style lang="scss">
	.markup-page {
		display: flex;

		gap: 20px;

		margin: 20px;

		height: 95%;

		.markup-view-container {
			display: flex;
			flex-direction: column;
			align-items: flex-start;

			gap: 20px;

			.markup-controls {
				display: flex;

				justify-content: space-between;
				align-items: center;

				width: 100%;
				gap: 12px;

				&__buttons {
					display: flex;

					gap: 16px;
				}
			}

			.canvas-container {
				box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
			}
		}

		.markup-storage {
			border-radius: 6px;

			width: 100%;

			box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
		}
	}
</style>
