<script lang="ts">
    import { AuthApi } from '$api/index';
	import { Button, Input } from '$lib/ui-components';
    import { authUser } from '$stores/index';
	import { navigate } from '$utils';

    const authApi = new AuthApi();
    let login: string;
    let password: string;

    const onSignIn = () => {
        (async () => {
            const response = await authApi.signIn({login, password});
            if(response !== null) {
                authUser.setUser(response.Result); 
                authUser.setUserInLocalStorage(response.Result);
                navigate("");
            }
        })();
    };

    const onSignUp = () => {
        (async () => {
            const response = await authApi.signUp({login, password});
            if(response !== null) {
                authUser.setUser(response.Result); 
                authUser.setUserInLocalStorage(response.Result);
                navigate("");
            }
        })();
    };
</script>

<section class="sign-in-page">
    <form class="container">
        <Input bind:value={login} />
        <Input bind:value={password} /> 
        <div class="container__buttons">
            <Button 
                type={"submit"}
                on:click={onSignIn}
            >
                Sign in
            </Button>
            <Button 
                type={"submit"}
                on:click={onSignUp}
            >
                Sign up
            </Button>
        </div>
    </form>
</section>