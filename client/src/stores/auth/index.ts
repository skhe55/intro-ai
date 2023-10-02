import { writable } from "svelte/store"
import type { TUser } from "$stores/types";
import { browser } from "$app/environment";

const createAuthStore = () => {
    const { subscribe, set } = writable({username: '', token: ''});

    return {
        subscribe,
        setUser: (payload: TUser) => set(payload),
        setUserInLocalStorage: (payload: Omit<TUser, "username"> & {expires_at: string}) => {
            if(browser) {
                window.localStorage.setItem("token", payload.token);
                if(payload.expires_at) {
                   window.localStorage.setItem("expires_at", new Date(payload.expires_at).toString())  
                } else {
                    window.localStorage.removeItem("expires_at");
                }
             }
        },
        get expiresAt() { return browser ? window.localStorage.getItem("expires_at") : ""; },
    };
};

export const authUser = createAuthStore();