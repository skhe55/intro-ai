<script lang="ts">
	import { onMount } from "svelte";
	import { ProjectApi } from "$api/index";
	import type { TProject } from "$api/types";

    let projectApi = new ProjectApi();
    let projects: TProject[] = [];

    onMount(() => {
        (async () => {
            const response = await projectApi.getProjects();
            if(response) {
                projects = [...projects, ...response.Result]
            }
        })();
    })
</script>

<div class="view">
    {#each projects as project (project.id)}
        <p>{project.name}</p>
    {/each}
</div>

<style lang="scss">
    .view {
        display: flex;
        flex-direction: column;

        height: 900px;
    }
</style>