<script lang="ts">
	import { afterUpdate, onMount } from "svelte";
  import { Stage, Layer, Rect, Image, type KonvaMouseEvent, Label, Tag, Text } from 'svelte-konva';
	import { AnnotationApi, ImageApi, LabelApi } from "$api/index";
	import type { TImage, TLabel } from "$api/types";
	import { DEFAULT_API_PATH, initialSquareCoordinates } from "$constants/index";
	import { Alert, Badge, Button, GradientButton, Input, Modal, Select, Label as TextLabel, Toast } from "flowbite-svelte";
	import type { TSquare } from "$lib/types";
	import Square from "$lib/shapes/square.svelte";
	import { getPointsUpperRightCorner, mappingAnnotationsToSquare } from "$lib/utils";
	import { ExclamationCircleSolid } from "flowbite-svelte-icons";
	  
  const STAGE_WIDTH = 600;
  const STAGE_HEIGHT = 600;

  let imageWidth: number;
  let imageHeight: number;

  let uuid = crypto.randomUUID();
  const imagesApi = new ImageApi();
  const labelApi = new LabelApi();
  const annotationApi = new AnnotationApi();

  let imageId: string | undefined = '';
  let currentImage: TImage;
  let imageElem: HTMLOrSVGImageElement;

  let labelName: string = '';
  let labels: TLabel[] = [];
  let selectedLabel: {id: string, name: string} | undefined;
  let isOpenModal: boolean = false;

  let isMarkupMode: boolean = false;
  let isWatchMode: boolean = false;
  let isEditMode: boolean = false;
  let squares: TSquare[] = [];
  let squareCoords: Omit<TSquare, 'id' | 'label_name' | 'restored'> = initialSquareCoordinates;

  let labelConfig = {
    x: 0,
    y: 0,
    opacity: 0.8,
    visible: false,
  };

  let labelTextConfig = {
    text: "",
    fontSize: 18,
    padding: 5,
    fill: "white",
  };

  let isShowToast: {msg: string, f: boolean, color: "green" | "orange" | "red"} = {msg: '', f: false, color: "green"};

  const onShowToast = (msg: string, f: boolean, color: "green" | "orange" | "red") => {
      isShowToast = {
          msg: msg,
          f: f,
          color: color,
      };
  };

  const onOpenModal = () => {
      isOpenModal = true;
  };
  
  const onCreateLabel = () => {
      (async () => {
          if(currentImage.id) {
            const response = await labelApi.createLabel({image_id: currentImage.id, name: labelName});
            if(response && response.Status == "OK") {
                const labelResponse = await labelApi.getLabelsByImageId(currentImage.id);
                if(labelResponse && labelResponse.Status === "OK") {
                  labels = [...labelResponse.Result];
                }
                onShowToast(`Successful created label!`, true, "green");
                setTimeout(() => {
                    onShowToast("", false, "green");
                }, 5000);
            }
            else {
                onShowToast(`Error occured while we creating label!`, true, "red");
                setTimeout(() => {
                    onShowToast("", false, "red");
                }, 5000);
            }
          }
      })();
  };

  const onCreateAnnotation = (labelId: string, coordinates: number[][]) => {
    (async () => {
      const response = await annotationApi.createAnnotation({label_id: labelId, image_id: currentImage.id, coordinates: coordinates});
      if(response && response.Status === "OK") {
        onShowToast(`Successful created annotation!`, true, "green");
        setTimeout(() => {
            onShowToast("", false, "green");
        }, 5000);
      } else {
        onShowToast(`Error occured while we creating annotation!`, true, "red");
        setTimeout(() => {
            onShowToast("", false, "red");
        }, 5000);
      }
    })();
  };

  const onDeleteAnnotation = (annotationId: string) => {
    (async () => {
      const response = await annotationApi.deleteAnnotation(annotationId);
      if(response && response.Status === "OK") {
        onShowToast(`Successful deleted annotation!`, true, "green");
        setTimeout(() => {
          onShowToast("", false, "green");
        }, 5000);
      } else {
        onShowToast(`Error occured while we deleting annotation!`, true, "red");
        setTimeout(() => {
          onShowToast("", false, "red");
        }, 5000);
      }
    })();
  };

  const onStartMarking = (e: KonvaMouseEvent) => {
    if(!selectedLabel || !selectedLabel.id) {
      onShowToast(`You don't choise a label, before starting mark it up, needed select label!`, true, "orange");
      setTimeout(() => {
        onShowToast("", false, "orange");
      }, 5000);
    } else {
      const { evt } = e.detail;
      if (isMarkupMode) {
        squares = [...squares, { ...squareCoords, id: uuid, label_name: selectedLabel.name, restored: false }];
        isMarkupMode = false;
        onCreateAnnotation(selectedLabel.id, [squareCoords.x_top, squareCoords.x_bottom, squareCoords.y_top, squareCoords.y_bottom]);
      } else if (!isEditMode && !isMarkupMode && !isWatchMode) {
        squareCoords = {
          x_top: [evt.offsetX, evt.offsetY],
          x_bottom: [evt.offsetX, evt.offsetY],
          y_top: [evt.offsetX, evt.offsetY],
          y_bottom: [evt.offsetX, evt.offsetY]
        };
        isMarkupMode = true;
      }
    }
  };

  const onMouseMove = (e: KonvaMouseEvent) => {
		const { evt } = e.detail;
    const imgPosition = e.detail.target.getRelativePointerPosition() as any;

		if (isMarkupMode) {
			squareCoords = {
				x_top: [squareCoords.x_top[0], squareCoords.x_top[1]],
				x_bottom: [squareCoords.x_top[0], evt.offsetY + imgPosition.y],
				y_top: [evt.offsetX + imgPosition.x, squareCoords.x_top[1]],
				y_bottom: [evt.offsetX + imgPosition.x, evt.offsetY + imgPosition.y]
			};
		}
	};

  const onRemoveItem = (_: KonvaMouseEvent, id: string, restored: boolean) => {
		squares = squares.filter((item) => item.id !== id);
    if(restored) {
      onDeleteAnnotation(id)
    };
	};

  const onMouseEnter = (e: KonvaMouseEvent) => {
    const konvaEvent = e.detail;

    let hoveredElementPos = konvaEvent.target.getPosition();
    let hoveredElementName = konvaEvent.target.attrs.name;

    labelConfig.x = hoveredElementPos.x;
    labelConfig.y = hoveredElementPos.y;

    labelTextConfig.text = hoveredElementName;

    labelConfig.visible = true;
  }

  const onMouseLeave = (_: KonvaMouseEvent) => {
    labelConfig.visible = false;
  };

  onMount(() => {
    window.addEventListener('keydown', (e) => {
			if (e.code === 'Escape') {
        isEditMode = false;
        isWatchMode = false;
				isMarkupMode = false;
			} else if (e.code === 'KeyW') {
				isEditMode = false;
        isWatchMode = true;
        isMarkupMode = false;
			} else if (e.code === 'KeyE') {
        isEditMode = true;
        isWatchMode = false;
        isMarkupMode = false;
      } 
		});

    (async () => {
        imageId = window.location.pathname.split("/").at(-1);
        if(imageId) {
          const response = await imagesApi.getImageById(imageId);  
          if(response && response.Status === "OK") {
            currentImage = {...response.Result};
            const labelsResponse = await labelApi.getLabelsByImageId(currentImage.id);
            if(labelsResponse && labelsResponse.Status === "OK") {
              labels = [...labelsResponse.Result];
            }
            const annotationsResponse = await annotationApi.getAnnotationByImageId(currentImage.id);
            if (annotationsResponse && annotationsResponse.Status === "OK") {
              squares = mappingAnnotationsToSquare(annotationsResponse.Result);
            }
            const img = document.createElement('img') as HTMLImageElement;
            img.src = `${DEFAULT_API_PATH}/static/${response.Result.projectId}/${response.Result.path_to_image}`;

            img.onload = () => {
              imageElem = img;
            };
          }
        }
    })();
  });

  afterUpdate(() => {
		uuid = crypto.randomUUID();
	});
</script>
<section class="markup-page">
  <Modal title="Create image" bind:open={isOpenModal} size={'xs'} autoclose>
    <form>
        <div>
            <TextLabel>Name</TextLabel>
            <Input bind:value={labelName} />
        </div>
    </form>
    <svelte:fragment slot="footer">
        <Button color="blue" disabled={labelName ? false : true} on:click={onCreateLabel}>Create</Button>
        <Button color="red">Cancel</Button>
    </svelte:fragment>
  </Modal>
  <Alert color="green">
    <p class="font-medium">
      Finally, we'll to mark up our image!
    </p>
    <p class="font-medium">
      If you don't have a labels, then we'll create a new labels!
    </p>
  </Alert>
  <div class="flex gap-3">
      <GradientButton color="tealToLime" on:click={() => window.history.back()}>
        Back to images
      </GradientButton >
      <GradientButton on:click={onOpenModal} color="tealToLime">
        Create label
      </GradientButton>
  </div>
  <div class="workspace-container">
    <Stage
      config={{ width: STAGE_WIDTH, height: STAGE_HEIGHT}}
      on:mousemove={onMouseMove}
      on:click={onStartMarking}
    >
      <Layer>
        <Image config={{ image: imageElem, width: imageWidth, height: imageHeight, draggable: true }} />
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
            on:mouseenter={onMouseEnter}
            on:mouseleave={onMouseLeave}
            rectConfig={{
              x: square.x_top[0],
              y: square.x_top[1],
              width: square.y_top[0] - square.x_top[0],
              height: square.y_bottom[1] - square.x_top[1],
              stroke: 'black',
              strokeWidth: 3,
              name: square.label_name,
            }}
            crossConfig={{
              points: getPointsUpperRightCorner(square),
              stroke: 'red',
              strokeWidth: 4
            }}
            crossOnClick={(e) => onRemoveItem(e, square.id, square.restored)}
            isShowCross={isEditMode}
          />
        {/each}
        <Label config={labelConfig}>
          <Tag
                config={{
                    fill: "black",
                    pointerDirection: "down",
                    pointerWidth: 10,
                    pointerHeight: 10,
                    lineJoin: "round",
                    shadowColor: "black",
                    shadowBlur: 10,
                    shadowOffsetX: 10,
                    shadowOffsetY: 10,
                    shadowOpacity: 0.5,
                }}
          />
            <Text config={labelTextConfig} />
        </Label>
      </Layer>
    </Stage>
    <div class="label-select-container">
      <Select bind:value={selectedLabel} placeholder={"Choose label"} class="w-80">
        {#each labels as label (label.id)}
          <option value={{id: label.id, name: label.name}}>{label.name}</option>
        {/each}
      </Select>
      <Alert color="green">
        <p class="font-medium">
          Hotkeys:
        </p>
        <p class="font-medium">
          Esc: reset all mods (use if during the process you realize that you do not want save annotation)
        </p>
        <p class="font-medium">
          W: enable watch mode
        </p>
        <p class="font-medium">
          E: enable edit mode
        </p>
      </Alert>
      {#if isShowToast.f}
        <Toast color={isShowToast.color}>
          <svelte:fragment slot="icon">
            <ExclamationCircleSolid class="w-5 h-5" />
            <span class="sr-only">Warning icon</span>
          </svelte:fragment>
          {isShowToast.msg}
        </Toast>
      {/if}
    </div>
  </div>
</section>

<style lang="scss">
  @use '../../../styles/lib/mixins.scss' as *;
	@use '../../../styles/lib/variables.scss' as *;

  .markup-page {
		display: flex;
    flex-direction: column;

    align-items: center;

		gap: 20px;

		margin: 20px;

		height: 95%;
  
    .workspace-container {
      display: flex;

      gap: 20px;

      .label-select-container {
        display: flex;
        flex-direction: column;

        gap: 10px;
      }
    }
	}
</style>
