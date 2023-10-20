<script lang="ts">
	import '$styles/app.scss';
	import { isTokenExpires, navigate } from '$utils';
	import { afterUpdate } from 'svelte';
	import { authUser } from '$stores/index';
	import { Button } from '$lib/ui-components';

	const onSignOut = () => {
		authUser.setUser({username: "", token: ""});
		authUser.setUserInLocalStorage({token: "", expires_at: ""});
		navigate("sign-in");
	};

	afterUpdate(() => {
		if(isTokenExpires(authUser.expiresAt) && window.location.pathname !== "/sign-in") {
			navigate("sign-in");
		} 
	});
</script>

<nav class="navbar">
	<div class="navbar__links-container">
		<a href="/markup">Markup</a>
	</div>
	<div class="navbar__user-container">
		<Button on:click={onSignOut}>
			Sign Out
		</Button>
	</div>
</nav>
<main>
	<slot />
</main>

<style lang="scss">
	@use '../styles/lib/mixins.scss' as *;
	@use '../styles/lib/variables.scss' as *;

	.navbar {
		display: flex;
		justify-content: space-between;
		align-items: center;

		background: #f6e4db;

		padding: 10px 0;

		max-height: 40px;
		height: 100%;

		&__links-container {
			display: flex;
			justify-content: center;
			align-items: center;

			gap: 25px;

			margin: 0 21px;
			a {
				@include text-large($black);

				text-decoration: none;

				background: #f6e4db;

				border-radius: 6px;

				transition: padding 1s ease-out, background 1s ease-in-out;

				&:hover {
					padding: 8px 16px;
					background: $link-hover-background-color;
				}
			}
		}

		&__user-container {
			display: flex;
			justify-content: center;
			align-items: center;


			gap: 20px;

			margin: 0 21px;
		}
	}

	main {
		max-height: 89%;
		height: 100%;

		background: $secondary-background-color;

		margin: 20px;

		border-radius: 6px;
		border: 1px solid #eef2f1;
	}
</style>
