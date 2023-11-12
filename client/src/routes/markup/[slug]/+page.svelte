<script lang="ts">
	import { onMount } from "svelte";
	import { ImageApi } from "$api/index";
	import type { TImage } from "$api/types";

    const imagesApi = new ImageApi();

    let imageId: string | undefined = '';
    let currentImage: TImage;
    
    onMount(() => {
        (async () => {
            imageId = window.location.pathname.split("/").at(-1);
            if(imageId) {
              const response = await imagesApi.getImageById(imageId);  
              if(response && response.Status === "OK") {
                currentImage = {...response.Result};
              }
            }
        })();
    })

    $: {
        console.log(currentImage);
    }
</script>

