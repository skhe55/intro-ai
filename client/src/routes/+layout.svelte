<script lang="ts">
	import '$styles/app.scss';
</script>

<nav>
	<a href="/">Home</a>
	<a href="/markup">Markup</a>
</nav>
<main>
	<slot />
</main>

<style lang="scss">
	@use '../styles/lib/mixins.scss' as *;
	@use '../styles/lib/variables.scss' as *;

	nav {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 20px;

		background: #f6e4db;

		padding: 10px 0;

		max-height: 40px;
		height: 100%;

		a {
			@include text-large($black);

			text-decoration: none;

			background: #f6e4db;

			border-radius: 6px;
			transition: padding 1s ease-out, background 1s ease-in-out;

			&:hover {
				padding: 8px 16px;
				background: #a0cbe1;
			}
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
