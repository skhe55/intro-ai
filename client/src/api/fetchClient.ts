import { navigate } from "$utils";

export const customFetch = (url: string, options: RequestInit | any = {}, authOptions?: any): Promise<Response> => {
    const OPTIONS = {
        url: '',
        unauthorizedRedirect: '',
        storage: window.localStorage,
        tokenName: 'token'
    }

    let opts = Object.assign({}, OPTIONS, authOptions);

    options.headers = options.headers || {};
    options.headers['Authorization'] = `Bearer ${opts.storage.getItem(opts.tokenName)}`;

    const request = window.fetch(url, options)
        .then((d) => {
            if (d.status === 401) {
                navigate("sign-in");
                return d
            } else {
                return d
            }
        });

    return request;
}