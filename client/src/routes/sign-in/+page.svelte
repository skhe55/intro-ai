<script lang="ts">
    import { AuthApi } from '$api/index';
    import { authUser } from '$stores/index';
	import { navigate } from '$utils';
    import { Input, Label, Button } from 'flowbite-svelte';

    const authApi = new AuthApi();
    let login: string;
    let password: string;

    const onSignIn = () => {
        (async () => {
            const response = await authApi.signIn({login, password});
            if(response !== null) {
                authUser.setUser(response.Result); 
                authUser.setUserInLocalStorage(response.Result);
                navigate("root");
            }
        })();
    };

    const onSignUp = () => {
        (async () => {
            const response = await authApi.signUp({login, password});
            if(response !== null) {
                authUser.setUser(response.Result); 
                authUser.setUserInLocalStorage(response.Result);
                navigate("root");
            }
        })();
    };
</script>

<section class="sign-in-page">
    <form class="container">
        <Label>Login</Label>
        <Input bind:value={login} />
        <Label>Password</Label>
        <Input bind:value={password} type={"password"} /> 
        <div class="container__buttons">
            <Button 
                color="alternative"
                on:click={onSignIn}
            >
                Sign in
            </Button>
            <Button 
                color="alternative"
                on:click={onSignUp}
            >
                Sign up
            </Button>
        </div>
    </form>
</section>