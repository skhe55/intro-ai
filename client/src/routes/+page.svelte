<script lang="ts">
	import { onMount } from "svelte";
	import { navigate } from "$utils";

    onMount(() => {
		if(window.location.pathname === '/') {
			navigate('root')
		}
	});
</script>