<script lang="ts">
	import type { TProject } from "$api/types";
	import { ImageIcon } from "$assets/index";

    export let projects: TProject[] = [];

</script>

<div class="view">
    {#each projects as project (project.id)}
        <div class="project">
            <p>{project.name}</p>
        </div>
        {#if project.images}
            <div class="image-container">
                {#each project.images as image (image.id)}
                    <div class="image">
                    <ImageIcon />
                        <p>{image.filename}</p> 
                    </div>
                {/each}
            </div>
        {/if}
    {/each}
</div>

<style lang="scss">
    @use "../../../../styles/lib/variables.scss" as *;
    @use "../../../../styles/lib/mixins.scss" as *;

    .view {
        display: flex;
        flex-direction: column;

        gap: 10px;
        margin: 20px;

        .project {
            width: auto;

            padding: 8px 12px;

            background-color: $third-background-color;

            border-radius: 6px;

            p {
                color: $black;
                font-weight: 550;
                font-size: $text-large-size;
            }
        }

        .image-container {
            display: flex;
            flex-direction: column;

            gap: 10px;

            overflow: auto;

            max-height: 300px;

            &::-webkit-scrollbar {
                display: block;
            }
            
            &::-webkit-scrollbar {
                width: 7px;
                height: 56px;
            }
            &::-webkit-scrollbar-track {
                width: 12px;
            }

            &::-webkit-scrollbar-thumb {
                border: 3px solid $third-background-color;
                background: $third-background-color;
                border-radius: 6px;
            }

            .image {
                cursor: pointer;

                display: flex;
                justify-content: flex-start;
                align-items: center;

                margin-left: 14px;
                padding: 6px;

                gap: 10px;

                background-color: $fourth-background-color;
                border-radius: 6px;

                max-width: 90%;

                p {
                    @include text-primary($black, 500);
                }
            }
        }
    }
</style>