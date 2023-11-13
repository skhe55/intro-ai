<script lang="ts">
	import { onMount } from "svelte";
  import { Stage, Layer, Rect, Image, type KonvaMouseEvent, Label } from 'svelte-konva';
	import { ImageApi, LabelApi } from "$api/index";
	import type { TImage } from "$api/types";
	import { DEFAULT_API_PATH } from "$constants/index";
	import { Alert, Button, GradientButton, Input, Modal, Label as TextLabel } from "flowbite-svelte";

    const imagesApi = new ImageApi();
    const labelApi = new LabelApi();

    let imageId: string | undefined = '';
    let currentImage: TImage;
    let imageElem: HTMLOrSVGImageElement;

    let labelName: string = '';
    let isOpenModal: boolean = false;

    let isShowToast: {msg: string, f: boolean, type: "error" | "success"} = {msg: '', f: false, type: "success"};

    const onShowToast = (msg: string, f: boolean, type: "error" | "success") => {
        isShowToast = {
            msg: msg,
            f: f,
            type,
        };
    };

    const onOpenModal = () => {
        isOpenModal = true;
    };
    
    const onCreateLabel = () => {
        (async () => {
            if(currentImage.projectId) {
              const response = await labelApi.createLabel({project_id: currentImage.projectId, name: labelName});
              if(response && response.Status == "OK") {
                const test = await labelApi.getLabelsByProjectId(currentImage.projectId);
                  onShowToast(`Successful created label!`, true, "success");
                  setTimeout(() => {
                      onShowToast("", false, "success");
                  }, 5000);
              }
              else {
                  onShowToast(`Error occured while we creating label!`, true, "error");
                  setTimeout(() => {
                      onShowToast("", false, "error");
                  }, 5000);
              }
            }
        })();
    };

    onMount(() => {
        (async () => {
            imageId = window.location.pathname.split("/").at(-1);
            if(imageId) {
              const response = await imagesApi.getImageById(imageId);  
              if(response && response.Status === "OK") {
                currentImage = {...response.Result};
                const img = document.createElement('img') as HTMLImageElement;
                img.src = `${DEFAULT_API_PATH}/static/${response.Result.projectId}/${response.Result.path_to_image}`;
                img.onload = () => {
                  imageElem = img;
                };
              }
            }
        })();
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
  <Stage
    config={{ width: 600, height: 600 }}
  >
    <Layer>
      <Image config={{ image: imageElem, width: 600, height: 600 }} />
    </Layer>
  </Stage>
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

		.markup-view-container {
			display: flex;
			flex-direction: column;
			align-items: flex-start;

			gap: 20px;

			.canvas-container {
				box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
			}
		}
	}
</style>
