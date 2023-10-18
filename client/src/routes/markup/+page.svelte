<script lang="ts">
	import { afterUpdate, onMount } from 'svelte';
    import { Stage, Layer, Rect, Image, type KonvaMouseEvent } from 'svelte-konva';

	import type { TSquare } from '$lib/types';
	import { initialSquareCoordinates } from '$constants/index';
	import Square from '$lib/shapes/square.svelte';
	import { getPointsUpperRightCorner } from '$lib/utils';
	import { Button, Modal, Input } from '$lib/ui-components';
	import { ProjectApi } from '$api/index';
	import { View } from './components';

	let projectApi = new ProjectApi();

	let uuid = crypto.randomUUID();

	let squares: TSquare[] = [];
	let squareCoords: Omit<TSquare, 'id'> = initialSquareCoordinates;

	let isMarkupMode: boolean = false;
	let isWatchMode: boolean = false;
	let isEditMode: boolean = false;

	let isShowCreateProjectModal: boolean = false;
	let nameOfCreatedProject: string = '';

	let image: HTMLOrSVGImageElement;

	const onToggleMode = () => {
		isEditMode = !isEditMode;
	};

	const onShowCreateProjectModal = () => {
		isShowCreateProjectModal = true;
	};

	const onCancelCreateProjectModal = () => {
		isShowCreateProjectModal = false;
		nameOfCreatedProject = '';
	};

	const onCreateProject = () => {
		(async () => {
			const response = await projectApi.createProject({name: nameOfCreatedProject});
			if(response) {
				onCancelCreateProjectModal();
			}
		})();
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

{#if isShowCreateProjectModal}
	<Modal>
		<div class="modal-create-project">
			<header class="modal-create-project-header">
				<h2>Create a project</h2>
			</header>
			<div class="modal-create-project__body">
				<div class="modal-create-project__inputs-container">
					<span class="label">Name</span>
					<Input className={"modal__inputs"} bind:value={nameOfCreatedProject} />
				</div>
			</div>
			<footer class="modal-create-project-footer">
				<div class="modal__buttons">
					<Button on:click={onCreateProject}>Approve</Button>
					<Button on:click={onCancelCreateProjectModal}>Cancel</Button>
				</div>
			</footer>
		</div>
	</Modal>
{/if}
<section class="markup-page">
	<div class="markup-view-container">
		<div class="markup-controls">
			<div class="markup-controls__buttons">
				<Button>Save</Button>
				<Button on:click={onShowCreateProjectModal}>Create a project</Button>
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
		<View />
	</div>
</section>

<style lang="scss">
	@use '../../styles/lib/mixins.scss' as *;
	@use '../../styles/lib/variables.scss' as *;

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

			overflow: auto;
		}
	}
	
	.modal-create-project {
		display: flex;
		flex-direction: column;

		align-items: center;
		justify-content: space-between;

		min-width: 400px;
		height: 200px;

		padding: 16px 0px;

		background-color: $white;

		border-radius: 6px;

		&__body {
			display: flex;
			flex-direction: column;

			gap: 16px;

			align-items: center;
			justify-content: center;
		}

		&__inputs-container {
			display: flex;
			flex-direction: column;

			.label {
				font-weight: 600;
				font-size: $text-large-size;
			}
		}
	}
</style>
