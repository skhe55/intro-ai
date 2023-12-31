<script lang="ts">
	import { onMount } from "svelte";
	import { AnnotationApi, ImageApi } from "$api/index";
	import type { TImage } from "$api/types";
	import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Button, Alert, GradientButton, Modal, Fileupload, Label, Helper, Toast, Input } from "flowbite-svelte";
	import { CheckCircleSolid, CloseCircleSolid } from "flowbite-svelte-icons";
	import { DEFAULT_API_PATH } from "$constants/index";
    import { format } from 'fecha';
	import { exportAnnotations } from "../../../lib/utils";

    let imagesApi = new ImageApi();
    let annotationApi = new AnnotationApi();

    let images: TImage[] = [];

    let uploadedImage: FileList | undefined;

    let projectId: string | undefined = '';
    let imageName: string = '';

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

    const onCreateImage = () => {
        (async () => {
            if(projectId && uploadedImage) {
                const createdImageResponse = await imagesApi.createImage({name: imageName, projectId: projectId});
                if (createdImageResponse && createdImageResponse.Status == "OK") {
                    const response = await imagesApi.uploadImage(createdImageResponse.Result, projectId, uploadedImage[0]);
                    if(response && response.Status == "OK") {
                        const response = await imagesApi.getImages(projectId);
                        if (response) {
                            images = [...response.Result];
                        }
                        onShowToast(`Successful upload image!`, true, "success");
                        setTimeout(() => {
                            onShowToast("", false, "success");
                        }, 5000);
                    }
                } else {
                    onShowToast(`Error occured while we uploading image!`, true, "error");
                    setTimeout(() => {
                        onShowToast("", false, "error");
                    }, 5000);
                }
            }
        })();
    };

    const onDeleteImage = (id: string, pathToImage: string) => {
        (async () => {
            if(projectId) {
                const response = await imagesApi.deleteImage(id, {project_id: projectId, path_to_image: pathToImage});
                if (response && response.Status === "OK") {
                    const response = await imagesApi.getImages(projectId); 
                    if(response) {
                        images = [...response.Result];
                    }
                    onShowToast(`Succesful deleted image!`, true, "success");
                    setTimeout(() => {
                        onShowToast("", false, "success");
                    }, 5000);
                } else {
                    onShowToast(`Error occured while we deleting image!`, true, "error");
                    setTimeout(() => {
                        onShowToast("", false, "error");
                    }, 5000);
                }
            }
        })();
    };

    const onExportAnnotations = (id: string) => {
        (async () => {
            const response = await annotationApi.getAnnotationByImageId(id);
            if(response && response.Status === "OK") {
                exportAnnotations(response.Result);
            }
        })();
    };

    onMount(() => {
        (async () => {
           projectId = window.location.pathname.split("/").at(-1);
            if(projectId) {
                const response = await imagesApi.getImages(projectId);
                if (response) {
                    images = [...images, ...response.Result];
                }
            }
        })();
    });
</script>

<section class="project-slug-page">
    <Modal title="Create image" bind:open={isOpenModal} size={'xs'} autoclose>
        <form>
            <div>
                <Label>Name</Label>
                <Input bind:value={imageName} />
            </div>
            <div>
                <Label for="picture" class="mb-2">Picture</Label>
                <Fileupload bind:files={uploadedImage} />
                <Helper helperClass="mt-2">Only JPG, PNG extensions</Helper>
            </div>
        </form>
        <svelte:fragment slot="footer">
            <Button color="blue" disabled={uploadedImage ? false : true} on:click={onCreateImage}>Create</Button>
            <Button color="red">Cancel</Button>
        </svelte:fragment>
    </Modal>
    <Alert color="green">
        <p class="font-medium">
            Now, time to choose image and start marking it up!
        </p>
        <p class="font-medium">
            If you don't have a image, then we'll create a new image!
        </p>
    </Alert>
    <div class="flex gap-3">
        <GradientButton color="tealToLime" on:click={() => window.history.back()}>
            Back to projects
        </GradientButton >
        <GradientButton color="tealToLime" on:click={onOpenModal}>
            Create image
        </GradientButton>
    </div>
    <div class="table-container">
        <Table divClass={"overflow-y-auto h-full mb-20"} striped={true}>
            <TableHead class="text-xs text-gray-700 uppercase bg-gray-100 dark:bg-gray-700 dark:text-gray-400">
                <TableHeadCell class="w-16">Id</TableHeadCell>
                <TableHeadCell class="w-32">Preview</TableHeadCell>
                <TableHeadCell class="w-32">Name</TableHeadCell>
                <TableHeadCell class="w-32">Created</TableHeadCell>
                <TableHeadCell class="w-16"></TableHeadCell>
                <TableHeadCell class="w-16"></TableHeadCell>
                <TableHeadCell class="w-16"></TableHeadCell>
            </TableHead>
            <TableBody>
                {#each images as image (image.id)}
                    <TableBodyRow>
                        <TableBodyCell>{image.id}</TableBodyCell>
                        <TableBodyCell>
                            <img class="preview-img" src={`${DEFAULT_API_PATH}/static/${projectId}/${image.path_to_image}`} alt={"image out of reach"} />
                        </TableBodyCell>
                        <TableBodyCell>{image.name}</TableBodyCell>
                        <TableBodyCell>{format(new Date(image.created_at), "MM.DD.YYYY")}</TableBodyCell>
                        <TableBodyCell>
                            <a href={`/markup/${image.id}`} class="font-medium text-primary-600 hover:underline dark:text-primary-500">
                                Open
                            </a>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Button on:click={() => onExportAnnotations(image.id)} color={"alternative"}>
                                Export annotations
                            </Button>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Button on:click={() => onDeleteImage(image.id, image.path_to_image)} color={"alternative"}>
                                Delete
                            </Button>
                        </TableBodyCell>
                    </TableBodyRow>
                {/each}
            </TableBody>
        </Table>
    </div>
    {#if isShowToast.f}
        <Toast color={isShowToast.type === "success" ? "green" : "red"}>
            <svelte:fragment slot="icon">
                {#if isShowToast.type === "success"}
                    <CheckCircleSolid class="w-5 h-5" />
                {:else}
                    <CloseCircleSolid class="w-5 h-5" />
                {/if}
            </svelte:fragment>
            <p class="font-medium">{isShowToast.msg}</p>
        </Toast>
    {/if}
</section>

<style lang="scss">
    @use '../../../styles/lib/mixins.scss' as *;
	@use '../../../styles/lib/variables.scss' as *;

    .project-slug-page {
		display: flex;
        flex-direction: column;

        align-items: center;

		gap: 20px;

		margin: 20px;

		height: 95%;

        .table-container {
            width: 100%;
            height: 95%;

            overflow: hidden;
        }

        .preview-img {
            width: 60px;
            height: 60px;
            object-fit: contain;
        }
	}
</style>
