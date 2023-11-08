<script lang="ts">
    import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Alert, GradientButton, Modal, Button, Label, Input, Toast  } from 'flowbite-svelte';
    import { CheckCircleSolid, CloseCircleSolid } from 'flowbite-svelte-icons';
    import { onMount } from 'svelte';
	import type { TProject } from '$api/types';
	import { ProjectApi } from '$api/index';

    let projects: TProject[] = [];
	let projectApi = new ProjectApi();
    
    let isOpenModal: boolean = false;

    let isShowToast: {msg: string, f: boolean, type: "error" | "success"} = {msg: '', f: false, type: "success"};

    let projectPayload: { name: string } = { name: '' };

    const onOpenModal = () => {
        isOpenModal = true;
    };

    const onShowToast = (msg: string, f: boolean, type: "error" | "success") => {
        isShowToast = {
            msg: msg,
            f: f,
            type,
        };
    };

    const onCreateProject = () => {
		(async () => {
			const response = await projectApi.createProject({name: projectPayload.name});
            if(response) {
                const response = await projectApi.getProjects();
                if(response) {
                    projects = [...response.Result];
                    onShowToast(`Successful created project with name: ${projectPayload.name}`, true, "success");
                    setTimeout(() => {
                        onShowToast("", false, "success");
                    }, 5000);
                }
            } else {
                onShowToast(`Error occured while we attempted create a project with name: ${projectPayload.name}`, true, "error");
                setTimeout(() => {
                    onShowToast("", false, "error");
                }, 5000);
            }
		})();
	};

    const onDeleteProject = (id: string) => {
        (async () => {
            const response = await projectApi.deleteProject(id);
            if(response) {
                const response = await projectApi.getProjects();
                if(response) {
                    projects = [...response.Result];
                    onShowToast(`Successful deleted project with id: ${id}`, true, "success");
                    setTimeout(() => {
                        onShowToast("", false, "success");
                    }, 5000);
                }
            } else {
                onShowToast(`Error occured while we attempted delete a project with id: ${id}`, true, "error");
                setTimeout(() => {
                    onShowToast("", false, "error");
                }, 5000);
            }
        })();
    };

    onMount(() => {
        (async () => {
            const response = await projectApi.getProjects();
            if(response) {
                projects = [...projects, ...response.Result]
            }
        })();
    });

   
</script>

<div class="root-page">
    <Modal title="Create a project" bind:open={isOpenModal} size={'xs'} autoclose>
        <form>
            <div>
                <Label for="project_name" class="mb-2">Project name</Label>
                <Input bind:value={projectPayload.name} type="text" id="project_name" required />
            </div>
        </form>
        <svelte:fragment slot="footer">
            <Button color="blue" disabled={projectPayload.name ? false : true} on:click={onCreateProject}>Create</Button>
            <Button color="red">Cancel</Button>
        </svelte:fragment>
    </Modal>
    <Alert color="green">
        <p class="font-medium">
            Welcome to root page!
            Choose project to open it!
        </p>
        <p class="font-medium">
            If you don't have a project, then we'll create a new project!
        </p>
    </Alert>
    <GradientButton color="tealToLime" on:click={onOpenModal}>
        Create project
    </GradientButton >
    <div class="table-container">
        <Table striped={true}>
            <TableHead class="text-xs text-gray-700 uppercase bg-gray-100 dark:bg-gray-700 dark:text-gray-400">
                <TableHeadCell>Project Name</TableHeadCell>
                <TableHeadCell class="w-16"></TableHeadCell>
                <TableHeadCell class="w-32"></TableHeadCell>
            </TableHead>
            <TableBody>
                {#each projects as project (project.id)}
                    <TableBodyRow>
                      <TableBodyCell>{project.name}</TableBodyCell>  
                      <TableBodyCell>
                        <a href="/markup" class="font-medium text-primary-600 hover:underline dark:text-primary-500">
                            Open
                        </a>
                      </TableBodyCell>
                      <TableBodyCell>
                        <Button color={"alternative"} on:click={() => onDeleteProject(project.id)}>
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
</div>

<style lang="scss">
    .root-page {
		display: flex;
        flex-direction: column;

        align-items: center;

		gap: 20px;

		margin: 20px;

		height: 95%;

        .table-container {
            width: 100%;
        }
	}
</style>



